package concurrent

import (
	"github.com/gorilla/websocket"
	"leeBlogCli/test/config"
	"log"
	"time"
)

type ResponseResult struct {
	Type     uint8       `json:"type"`
	Time     *int        `json:"-" `
	Code     int         `json:"code"`
	Data     interface{} `json:"data"`
	Files    interface{} `json:"files"`
	Markdown interface{} `json:"markdown"`
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
}

func (w *Writer) Run() {

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
