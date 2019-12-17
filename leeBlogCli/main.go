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
	http.HandleFunc(config.NewFile, handler.ReceivingFile)
	http.HandleFunc(config.DeleteFile, handler.DeleteFile)
	fmt.Printf("server start with http://localhost:%s\n", config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}