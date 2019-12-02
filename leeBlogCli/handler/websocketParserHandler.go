package handler

import (
	"encoding/binary"
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"leeBlogCli/concurrent"
	"leeBlogCli/parser"
	"log"
	"net/http"
	"strconv"
	"time"
)

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

func websocketLoop(conn *websocket.Conn, writer *concurrent.Writer) {
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
			t := time.Now().Nanosecond()
			uuid := uuid2.New()
			go func(t *int, uuid uuid2.UUID) {
				defer func() {
					if err := recover(); err != nil {
						log.Printf("text parse error %s", err)
					}
				}()
				var obj = ParamNewArticle{
					Type:  1,
					Files: nil,
				}

				if err := json.Unmarshal(p, &obj); err != nil {
					log.Println(err)
					errResult := concurrent.ResponseResult{
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
					if obj.Files != nil {
						//log.Printf("%v", obj.Files)
						for i := range obj.Files {
							obj.Files[i].ID = uuid2.New().ID()
							writer.FileServer.FileMap[obj.Files[i].ID] = &concurrent.FileInfo{
								ID:          obj.Files[i].ID,
								Name:        obj.Files[i].Name,
								ExtType:     obj.Files[i].ExtType,
								ServerFile:  nil,
								BufIOWriter: nil,
								ServerName:  strconv.FormatUint(uint64(obj.Files[i].ID), 10),
							}
						}
						result := concurrent.ResponseResult{Type: 2, Time: t, Code: 0, Files: obj.Files}
						writer.ResultChan <- &result
					}
				}
			}(&t, uuid)
		} else if messageType == websocket.BinaryMessage {
			fragment := concurrent.FileFragment{
				FileID:          binary.BigEndian.Uint32(p[0:4]),
				FragmentIndex:   binary.BigEndian.Uint16(p[4:6]),
				FileFragmentEnd: binary.BigEndian.Uint16(p[6:8]) == 1,
				FragmentData:    p[8:],
			}
			writer.FileServer.FileFragmentChan <- &fragment
		}
	}
}

func markdownParse(p string) concurrent.ResponseResult {
	result := concurrent.ResponseResult{}
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

func WebSocketReadMarkdownText(writer http.ResponseWriter, r *http.Request) {
	//header := http.Header{}
	//header.Add("Access-Control-Allow-Origin", "http://localhost:3000")
	conn, err := upgrade.Upgrade(writer, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	socketWriter := concurrent.Writer{
		Conn:       conn,
		ResultChan: make(chan *concurrent.ResponseResult),
		FileServer: concurrent.FileServer{
			FileMap:          make(map[uint32]*concurrent.FileInfo, 0),
			FileFragmentChan: make(chan *concurrent.FileFragment),
		},
	}
	websocketLoop(conn, &socketWriter)
}
