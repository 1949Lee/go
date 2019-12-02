package concurrent

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// 接收上传的文件
func ReceivingFile(writer http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	/**
	  此部分代码可以提取出来。*/
	//设置跨域的相应头CORS，CORS参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS
	writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	// 请求方发送请求时当请求为options时，直接返回200。
	if r.Method == "OPTIONS" {
		writer.WriteHeader(200)
		return
	}
	t := time.Now()
	//下面是接口真正的处理
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	err := r.ParseMultipartForm(64 << 10)
	if err != nil {
		log.Printf("ReceivingFile Handler when ParseMultipartForm Error:%v", err)
	}
	for _, v := range r.MultipartForm.File {
		for i := 0; i < len(v); i++ {
			file, err := v[i].Open()
			if err != nil {
				log.Printf("ReceivingFile Handler when FormFile Error:%v", err)
			}
			buffer := bufio.NewReader(file)
			f, err := os.Create(v[i].Filename)
			if err != nil {
				log.Printf("ReceivingFile Handler when os.Create Error:%v", err)
			}
			_, err = buffer.WriteTo(f)
			if err != nil {
				log.Printf("ReceivingFile Handler when buffer.WriteTo Error:%v", err)
			}
			_ = file.Close()
		}
	}

	result := ResponseResult{
		Type: 4,
		Code: 0,
		Data: "接收成功",
	}
	var b []byte
	if b, err = json.Marshal(result); err != nil {
		log.Printf("ReceivingFile Handler when json.Marshal Error:%v", err)
	}

	writer.Write([]byte(b))
	fmt.Println("app elapsed:", time.Since(t))
}
