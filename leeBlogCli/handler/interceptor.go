package handler

import (
	"log"
	"net/http"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

// Http拦截器
func HttpInterceptor(handler HttpHandler, header map[string]string) HttpHandler {
	return func(writer http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		for k := range header {
			writer.Header().Add(k, header[k])
		}
		// 请求方发送请求时当请求为options时，直接返回200。
		if r.Method == "OPTIONS" {
			writer.WriteHeader(200)
			return
		}
		//下面是接口真正的处理
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		handler(writer, r)
	}
}

// 普通接口拦截
func APIInterceptor(handler HttpHandler) HttpHandler {
	return HttpInterceptor(handler, map[string]string{
		"Access-Control-Allow-Methods": "POST, OPTIONS",
		"Access-Control-Allow-Headers": "POST, Content-Type",
	})
}

// 资源拦截
func ResourceInterceptor(handler HttpHandler) HttpHandler {
	// 需要加严格的验证。
	return HttpInterceptor(handler, map[string]string{
		"Access-Control-Allow-Methods": "GET",
	})
}
