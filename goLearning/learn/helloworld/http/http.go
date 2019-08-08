package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// 建立一个请求
	request, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)

	// 设置请求头
	//request.Header.Add("User-Agent","XXXXXX")

	//建立客户端，
	client := http.Client{}

	//用建立的客户端发送请求
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer  response.Body.Close()
	bytes, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",bytes)
}
