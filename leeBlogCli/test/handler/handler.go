package handler

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"leeBlogCli/test/concurrent"
	"leeBlogCli/test/fileServer"
	"leeBlogCli/test/parser"
	"log"
	"net/http"
	"os"
	"time"
)

type ParamNewArticle struct {
	Text string
	File fileServer.File
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		if r.Header.Get("Origin") == "http://localhost:3000" {
			return true
		}
		return false
	},
}

func ReadMarkdownText(writer http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	/**
	  此部分代码可以提取出来。*/
	//设置跨域的相应头CORS，CORS参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS
	writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	// 请求方发送请求时当请求为options时，直接返回200。
	if r.Method == "OPTIONS" {
		writer.WriteHeader(200)
		return
	}
	t := time.Now()
	//下面是接口真正的处理
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	var param ParamNewArticle
	result := concurrent.ResponseResult{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err := json.Unmarshal(body, &param)
	if err != nil {
		writer.Write([]byte("error params"))
	}
	result.Code = 0

	dataList, html := parser.MarkdownParse(param.Text)
	result.Data = struct {
		Text         string              `json:"text"`
		List         []parser.TokenSlice `json:"list"`
		MarkDownHtml string              `json:"html"`
	}{
		Text:         "success",
		List:         dataList,
		MarkDownHtml: html,
	}
	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error:", err)
	}
	writer.Write([]byte(b))
	fmt.Println("app elapsed:", time.Since(t))
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
			go func(t *int) {
				var obj ParamNewArticle

				if err := json.Unmarshal(p, &obj); err != nil {
					log.Println(err)
					errResult := concurrent.ResponseResult{
						Time: t,
						Code: 1,
						Data: "无法识别的参数",
					}
					writer.ResultChan <- &errResult
				}
				if obj.Text != "" { // 有text表示就是要转换markdown
					result := markdownParse(obj.Text)
					result.Time = t
					writer.ResultChan <- &result
				}
				if obj.File != (fileServer.File{}) {
					log.Printf("%v", obj.File)
					result := concurrent.ResponseResult{Time: t, Code: 0, Data: "收到了文件信息"}
					writer.ResultChan <- &result
				}
			}(&t)
		} else if messageType == websocket.BinaryMessage {
			log.Printf("%d", binary.BigEndian.Uint16(p[0:2]))
			log.Printf("%d", binary.BigEndian.Uint16(p[2:4]))
			f, err := os.Create("./test.jpg")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			bw := bufio.NewWriter(f)
			_, ok := bw.Write(p[4:])
			if ok != nil {
				log.Println(ok)
			}
			yes := bw.Flush()
			if yes != nil {
				log.Println(yes)
			}
		}
	}
}

func markdownParse(p string) concurrent.ResponseResult {
	result := concurrent.ResponseResult{}
	result.Code = 0
	dataList, html := parser.MarkdownParse(p)
	result.Markdown = struct {
		Text         string              `json:"text"`
		List         []parser.TokenSlice `json:"list"`
		MarkDownHtml string              `json:"html"`
	}{
		Text:         "success",
		List:         dataList,
		MarkDownHtml: html,
	}
	return result
}

func SocketReadMarkdownText(writer http.ResponseWriter, r *http.Request) {
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
	}
	websocketLoop(conn, &socketWriter)
}
