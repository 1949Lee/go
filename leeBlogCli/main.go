package main

import (
	"fmt"
	"leeBlogCli/concurrent"
	"leeBlogCli/config"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	lee := concurrent.Lee{}
	lee.Run()
	defer lee.Close()
	http.HandleFunc(config.WebsocketParserPath, lee.API.WebSocketReadMarkdownText)
	http.HandleFunc(config.NewFile, lee.API.APIInterceptor(lee.API.ReceivingFile))
	http.HandleFunc(config.DeleteFile, lee.API.APIInterceptor(lee.API.DeleteFile))
	http.HandleFunc(config.NewArticleID, lee.API.APIInterceptor(lee.API.NewArticleID))
	http.HandleFunc(config.GetArticleWithEditingInfo, lee.API.APIInterceptor(lee.API.GetArticleWithEditingInfo))
	http.HandleFunc(config.TagsGroupByCategory, lee.API.APIInterceptor(lee.API.GetTagsGroupByCategory))
	http.HandleFunc(config.NewTag, lee.API.APIInterceptor(lee.API.NewTag))
	http.HandleFunc(config.DeleteTag, lee.API.APIInterceptor(lee.API.DeleteTag))
	http.HandleFunc(config.NewCategory, lee.API.APIInterceptor(lee.API.NewCategory))
	http.HandleFunc(config.DeleteCategory, lee.API.APIInterceptor(lee.API.DeleteCategory))
	http.HandleFunc(config.FileResource, lee.API.ResourceInterceptor(lee.API.FileResource))
	fmt.Printf("server start with http://localhost:%s\n", config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
