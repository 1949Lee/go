package main

import (
	"fmt"
	"leeBlogCli/test/config"
	"leeBlogCli/test/handler"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/once", handler.ReadMarkdownText)
	http.HandleFunc(config.WebsocketParserPath, handler.WebSocketReadMarkdownText)
	fmt.Printf("server start with http://localhost:%s\n", config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
	//listener, err := net.Listen("tcp", ":1315")
	//if err!=nil{
	//    panic(err)
	//}
	//for {
	//    conn, err := listener.Accept()
	//    if err!=nil{
	//        log.Printf("accept error:%v",err)
	//        continue
	//    }
	//
	//    go func(conn net.Conn) {
	//        handler.SocketReadMarkdownText(conn)
	//    }(conn)
	//}
}
