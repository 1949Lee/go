package handler

import (
	"io/ioutil"
	"leeBlogCli/config"
	"log"
	"net/http"
	"os"
)

func FileResource(writer http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	/**
	  此部分代码可以提取出来。*/
	//设置跨域的相应头CORS，CORS参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS
	writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
	writer.Header().Add("Access-Control-Allow-Methods", "GET")

	// 需要加严格的验证。

	//下面是接口真正的处理
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	requestUrl := r.URL.Path
	filePath := requestUrl[len(config.FileResource):]
	file, err := os.Open(config.FilePath + filePath)
	defer file.Close()
	if err != nil {
		log.Println("static resource:", err)
		writer.WriteHeader(404)
	} else {
		bs, _ := ioutil.ReadAll(file)

		writer.Write(bs)
	}
}
