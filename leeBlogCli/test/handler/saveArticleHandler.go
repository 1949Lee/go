package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"leeBlogCli/test/concurrent"
	"leeBlogCli/test/parser"
	"log"
	"net/http"
	"time"
)

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
