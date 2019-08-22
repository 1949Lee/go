package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"leeBlogCli/test/parser"
	"net/http"
	"strings"
	"time"
)

type ParamNewArticle struct {
	Text string
}

type ResponseResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ReadMarkdownText(writer http.ResponseWriter, r *http.Request) {

	/**
	  此部分代码可以提取出来。*/
	//设置跨域的相应头CORS，CORS参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS
	writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:63342")
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
	var list []string
	//scanner := bufio.NewScanner(strings.NewReader(param.Text))
	//for scanner.Scan() {
	//	list = append(list, scanner.Text())
	//}

	// 这种split的方法比bufio那种读取块100-500微秒。
	list = strings.Split(param.Text, "\n")
	line := parser.Line{Origin: []rune(list[0]), Tokens: []parser.Token{}}
	line.Parse()
	result.Data = struct {
		Text string         `json:"text"`
		List []parser.Token `json:"list"`
	}{
		Text: "success",
		List: line.Tokens,
	}
	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error:", err)
	}
	writer.Write([]byte(b))
	fmt.Println("app elapsed:", time.Since(t))
}
