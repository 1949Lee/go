package main

import (
	"fmt"
	"leeBlogCli/test/handler"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	port := "1314"
	http.HandleFunc("/once", handler.ReadMarkdownText)
	http.HandleFunc("/ws", handler.SocketReadMarkdownText)
	fmt.Printf("server start with http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
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
