package main

import (
	"fmt"
	"leeBlogCli/concurrent"
	"leeBlogCli/config"
	"leeBlogCli/dao"
	"leeBlogCli/handler"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	db := dao.DBServer{}
	db.Open()
	defer db.Close()
	http.HandleFunc("/once", handler.ReadMarkdownText)
	http.HandleFunc(config.WebsocketParserPath, handler.WebSocketReadMarkdownText)
	http.HandleFunc(config.NewFile, concurrent.ReceivingFile)
	http.HandleFunc(config.DeleteFile, concurrent.DeleteFile)
	fmt.Printf("server start with http://localhost:%s\n", config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
