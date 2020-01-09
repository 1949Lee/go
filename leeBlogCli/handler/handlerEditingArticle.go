package handler

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"leeBlogCli/parser"
	"leeBlogCli/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type WebSocketWriter struct {
	messageID  int64
	Conn       *websocket.Conn
	ResultChan chan *definition.ResponseResult
	FileServer FileServer
}

func (w *WebSocketWriter) Run() {

	// 运行文件服务器
	w.FileServer.Run(w)

	go func() {
		var resultQueue definition.ResponseResultQueue
		ticker := time.Tick(config.WebsocketTickerDuration)
		// 每隔一段时间，发送这段时间内前端最后一次请求的编译结果。
		go func(ticker <-chan time.Time) {
			for range ticker {
				result := resultQueue.Max()
				if result != nil {
					if err := w.Conn.WriteJSON(result); err != nil {
						log.Printf("write err:%v", err)
					}
					resultQueue = definition.ResponseResultQueue{}
				}
			}
		}(ticker)
		for result := range w.ResultChan {
			if result.Type != 1 {
				if err := w.Conn.WriteJSON(result); err != nil {
					log.Printf("write err:%v", err)
				}
			} else {
				resultQueue = append(resultQueue, result)
			}
		}
	}()
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		if r.Header.Get("Origin") == "http://localhost:8080" {
			return true
		}
		return false
	},
}

func websocketLoop(conn *websocket.Conn, writer *WebSocketWriter) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("websocketLoop error %s", err)
			//if err = conn.WriteJSON(conn.WriteJSON(ResponseResult{Code: 1})); err != nil {
			//	log.Printf("write err:%v", err)
			//	//return
			//}
			websocketLoop(conn, writer)
		}
	}()
	writer.Run()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			_, ok := err.(*websocket.CloseError)
			if ok {
				break
			}
			log.Printf("receive err:%v", err)
			continue
		}
		if messageType == websocket.TextMessage {
			t := writer.messageID
			writer.messageID = writer.messageID + 1
			//t := int64(time.Now().Nanosecond())
			uuid := uuid2.New()
			go func(t *int64, uuid uuid2.UUID) {
				defer func() {
					if err := recover(); err != nil {
						log.Printf("text parse error %s", err)
					}
				}()
				var obj = definition.ParamEditingArticle{
					ArticleID: -1,
					Type:      1,
					Files:     nil,
				}

				if err := json.Unmarshal(p, &obj); err != nil {
					log.Println(err)
					errResult := definition.ResponseResult{
						Time: t,
						Code: 1,
						Data: "无法识别的参数",
					}
					writer.ResultChan <- &errResult
				}
				switch obj.Type {
				case 1:
					result := markdownParse(obj.Text)
					result.Time = t
					result.Type = 1
					writer.ResultChan <- &result
				case 2:
					if obj.Files != nil && obj.ArticleID != -1 {
						//log.Printf("%v", obj.Files)
						if err := os.MkdirAll(utils.GetFilePath(obj.ArticleID), os.ModePerm); err != nil {
							//log.Printf("ReceivingFile Handler when os.MkdirAll Error:%v", err)
							//result.Code = 1
							//result.Data = "服务器保存文件失败"
						}
						for i := range obj.Files {
							obj.Files[i].ID = uuid2.New().ID()
							writer.FileServer.FileMap[obj.Files[i].ID] = &definition.FileInfo{
								ID:          obj.Files[i].ID,
								Name:        obj.Files[i].Name,
								ExtType:     obj.Files[i].ExtType,
								ServerFile:  nil,
								BufIOWriter: nil,
								ServerName:  strconv.FormatUint(uint64(obj.Files[i].ID), 10),
								ArticleID:   obj.ArticleID,
							}
						}
						result := definition.ResponseResult{Type: 2, Time: t, Code: 0, Files: obj.Files}
						writer.ResultChan <- &result
					}
				}
			}(&t, uuid)
		} else if messageType == websocket.BinaryMessage {
			fragment := definition.FileFragment{
				FileID:          binary.BigEndian.Uint32(p[0:4]),
				FragmentIndex:   binary.BigEndian.Uint16(p[4:6]),
				FileFragmentEnd: binary.BigEndian.Uint16(p[6:8]) == 1,
				FragmentData:    p[8:],
			}
			writer.FileServer.FileFragmentChan <- &fragment
		}
	}
}

func markdownParse(p string) definition.ResponseResult {
	result := definition.ResponseResult{}
	result.Code = 0
	dataList, _ := parser.MarkdownParse(p)
	result.Markdown = struct {
		Text string              `json:"text"`
		List []parser.TokenSlice `json:"list"`
		//MarkDownHtml string              `json:"html"`
	}{
		Text: "success",
		List: dataList,
		//MarkDownHtml: html,
	}
	return result
}

func (api *API) WebSocketReadMarkdownText(writer http.ResponseWriter, r *http.Request) {
	//header := http.Header{}
	//header.Add("Access-Control-Allow-Origin", "http://localhost:3000")
	conn, err := upgrade.Upgrade(writer, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	socketWriter := WebSocketWriter{
		Conn:       conn,
		ResultChan: make(chan *definition.ResponseResult),
		FileServer: FileServer{
			FileMap:          make(map[uint32]*definition.FileInfo, 0),
			FileFragmentChan: make(chan *definition.FileFragment),
		},
	}
	websocketLoop(conn, &socketWriter)
}

// TODO 上传的图片需要处理，渐进式显示图片，图片压缩。https://godoc.org/gopkg.in/gographics/imagick.v2/imagick、https://github.com/gographics/imagick、https://godoc.org/gopkg.in/gographics/imagick.v2/imagick、https://godoc.org/gopkg.in/gographics/imagick.v3/imagick

// 前两个字节（16位）表示每个文件的第几个分片（最大共65536个分片）；后两个字节表示每个分片的大小（单位kb，默认64kb）。所以单个文件最大64kb * 2^16=4G
// 文件服务器：上传下载相关。
type FileServer struct {
	// 处理中的文件映射
	FileMap definition.FileMap

	// 文件片段管道
	FileFragmentChan chan *definition.FileFragment
}

// 文件服务器
func (s *FileServer) Run(w *WebSocketWriter) {
	go func() {
		//var fragmentQueue FileFragmentQueue
		for fragment := range s.FileFragmentChan {
			//if result.Type != 1 {
			//    if err := w.Conn.WriteJSON(result); err != nil {
			//        log.Printf("write err:%v", err)
			//    }
			//} else {
			//    fragmentQueue = append(fragmentQueue, fragment)
			//}

			info := s.FileMap[fragment.FileID]

			go func(fileInfo *definition.FileInfo) {
				var builder strings.Builder
				builder.WriteString(config.FilePath)
				builder.WriteString(strconv.Itoa(info.ArticleID))
				builder.WriteString("/")
				builder.WriteString(fileInfo.ServerName)
				builder.WriteString(".lee")
				temName := builder.String()
				builder.Reset()
				builder.WriteString(config.FilePath)
				builder.WriteString(strconv.Itoa(info.ArticleID))
				builder.WriteString("/")
				builder.WriteString(fileInfo.ServerName)
				builder.WriteString(".")
				builder.WriteString(fileInfo.ExtType)
				serverName := builder.String()
				if fileInfo.ServerFile == nil {
					//filepath.Dir()
					f, err := os.Create(temName)
					if err != nil {
						log.Printf("create file fail, the err:%v", err)
						w.ResultChan <- &definition.ResponseResult{Type: 3, Code: definition.FileStatus.InitFail}
					}
					fileInfo.ServerFile = f
					fileInfo.BufIOWriter = bufio.NewWriter(f)
				}
				_, ok := fileInfo.BufIOWriter.Write(fragment.FragmentData)
				if ok != nil {
					log.Printf("Write file fagment error, fileID:%d fileName:%s fragmentIndex:%d error is: %v \n", fileInfo.ID, fileInfo.Name, fragment.FragmentIndex, ok)
				}
				yes := fileInfo.BufIOWriter.Flush()
				if yes != nil {
					log.Println(yes)
				}
				// 表示文件数据到了最后。
				if fragment.FileFragmentEnd {
					err := os.Rename(temName, serverName)
					if err != nil {
						log.Printf("Rename file(%s) error, error is: %v \n", temName, err)
						delete(s.FileMap, fileInfo.ID)
						fileInfo.ServerFile.Close()
						if err := os.Remove(temName); err != nil {
							log.Printf("Remove file(%s) error, error is: %v \n", temName, err)
						}
						w.ResultChan <- &definition.ResponseResult{Type: 3, Code: definition.FileStatus.ProcessFail}
					}
					fileInfo.ServerFile.Close()
				}
			}(info)
		}
	}()
}
