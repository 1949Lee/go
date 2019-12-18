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
	http.HandleFunc(config.FileResource, lee.API.ResourceInterceptor(lee.API.FileResource))
	fmt.Printf("server start with http://localhost:%s\n", config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
