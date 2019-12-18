package main

import (
	"fmt"
	"leeBlogCli/concurrent"
	"leeBlogCli/config"
	"leeBlogCli/handler"
	websocketHandler "leeBlogCli/handler/websocket"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	blog := concurrent.Blog{}
	blog.Run()
	defer blog.Close()
	http.HandleFunc(config.WebsocketParserPath, websocketHandler.WebSocketReadMarkdownText)
	http.HandleFunc(config.NewFile, handler.APIInterceptor(handler.ReceivingFile))
	http.HandleFunc(config.DeleteFile, handler.APIInterceptor(handler.DeleteFile))
	http.HandleFunc(config.FileResource, handler.ResourceInterceptor(handler.FileResource))
	fmt.Printf("server start with http://localhost:%s\n", config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
