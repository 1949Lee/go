package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"leeBlogCli/test/parser"
	"log"
	"net/http"
	"time"
)

type ParamNewArticle struct {
	Text string
}

type ResponseResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{
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
	result := ResponseResult{}
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

func SocketReadMarkdownText(writer http.ResponseWriter, r *http.Request) {
	//header := http.Header{}
	//header.Add("Access-Control-Allow-Origin", "http://localhost:3000")
	conn, err := upgrader.Upgrade(writer, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		var param ParamNewArticle
		result := ResponseResult{}
		err := conn.ReadJSON(&param)

		if err != nil {
			//if err = {
			//    break
			//}
			_, ok := err.(*websocket.CloseError)
			if ok {
				break
			}
			log.Printf("receive err:%v", err)
			continue
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
		if err = conn.WriteJSON(result); err != nil {
			log.Printf("receive err:%v", err)
			break
		}
	}

}
