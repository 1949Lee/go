package concurrent

import (
	"github.com/gorilla/websocket"
	"log"
)

type ResponseResult struct {
	Code     int         `json:"code"`
	Data     interface{} `json:"data"`
	Markdown interface{} `json:"markdown"`
}

type Writer struct {
	Conn       *websocket.Conn
	ResultChan chan *ResponseResult
}

func (w *Writer) Run() {

	go func() {
		for result := range w.ResultChan {
			if err := w.Conn.WriteJSON(result); err != nil {
				log.Printf("write err:%v", err)
			}
		}
	}()
}
