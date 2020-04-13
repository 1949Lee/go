package handler

import (
	"encoding/json"
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"log"
	"net/http"
	"strings"
)

type APIResponseWriter struct {
	writer http.ResponseWriter
}

func (a *APIResponseWriter) Send(result interface{}) (int, error) {
	var (
		b   []byte
		err error
	)
	if b, err = json.Marshal(result); err != nil {
		log.Printf("Send API result when json.Marshal Error:%v", err)
	}

	return a.writer.Write(b)
}

func (a *APIResponseWriter) Write(b []byte) (int, error) {
	return a.writer.Write(b)
}

type HttpHandler func(http.ResponseWriter, *http.Request)
type APIHandler func(*APIResponseWriter, *http.Request)

// Http拦截器
func (api *API) HttpInterceptor(handler APIHandler, header map[string]string, options definition.InterceptorOptions) HttpHandler {
	return func(writer http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		origin := r.Header.Get("Origin")
		if !strings.HasSuffix(origin, config.LegalOriginURL) {
			writer.WriteHeader(403)
			return
		}
		writer.Header().Add("Access-Control-Allow-Origin", origin)
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
		if options.CheckLogin {
			// 验证登录
			loginStatus := api.CheckLogin(r)
			if !loginStatus {
				//用户未登录
				result := definition.APIResult{
					Code: definition.ResponseServerCode.NotLogin,
					Data: "用户未登录",
				}
				tempWriter := APIResponseWriter{
					writer: writer,
				}
				_, _ = tempWriter.Send(result)
				return
			}
		}
		apiWriter := APIResponseWriter{
			writer: writer,
		}
		handler(&apiWriter, r)
	}
}

// 普通接口拦截
func (api *API) APIInterceptor(handler APIHandler, options definition.InterceptorOptions) HttpHandler {
	return api.HttpInterceptor(handler, map[string]string{
		"Access-Control-Allow-Methods": "POST, OPTIONS",
		"Access-Control-Allow-Headers": "POST, Content-Type,leeKey,leeToken",
	}, options)
}

// 资源拦截
func (api *API) ResourceInterceptor(handler APIHandler) HttpHandler {
	// 需要加严格的验证。
	return func(writer http.ResponseWriter, r *http.Request) {
		// 请求方发送请求时当请求为options时，直接返回200。
		if r.Method == "Get" {
			writer.WriteHeader(403)
			return
		}
		apiWriter := APIResponseWriter{
			writer: writer,
		}
		handler(&apiWriter, r)
	}
}
