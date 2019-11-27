package concurrent

import (
	"github.com/gorilla/websocket"
	"leeBlogCli/test/config"
	"log"
	"time"
)

type ResponseCodeType int

type ResponseResult struct {
	// 1表示markdown相关，2 表示文件准备相关，3表示文件上传相关。
	Type uint8 `json:"type"`
	Time *int  `json:"-" `

	// code码开头第一位表示type类型的值。如文件上传相关则为形如：3XX
	Code     ResponseCodeType `json:"code"`
	Data     interface{}      `json:"data"`
	Files    interface{}      `json:"files"`
	Markdown interface{}      `json:"markdown"`
}

type ResponseResultQueue []*ResponseResult

func (q *ResponseResultQueue) Max() *ResponseResult {
	list := *q
	var result *ResponseResult
	if len(list) > 0 {
		for i := len(list) - 1; i >= 0; i-- {
			if result == nil {
				result = list[i]
			}
			if *(result.Time) < *(list[i].Time) {
				result = list[i]
			}
		}
	}
	return result
}

type Writer struct {
	Conn       *websocket.Conn
	ResultChan chan *ResponseResult
	FileServer FileServer
}

func (w *Writer) Run() {

	// 运行文件服务器
	w.FileServer.Run(w.Conn)

	go func() {
		var resultQueue ResponseResultQueue
		ticker := time.Tick(config.WebsocketTickerDuration)
		// 每隔一段时间，发送这段时间内前端最后一次请求的编译结果。
		go func(ticker <-chan time.Time) {
			for range ticker {
				result := resultQueue.Max()
				if result != nil {
					if err := w.Conn.WriteJSON(result); err != nil {
						log.Printf("write err:%v", err)
					} else {
						resultQueue = ResponseResultQueue{}
					}
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
