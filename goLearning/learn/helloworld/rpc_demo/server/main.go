package main

import (
	"goLearning/learn/helloworld/rpc_demo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpc_demo.DemoService{})
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1949")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}

		// 开一个协程去执行，否则会阻塞下一个调用rpc的人
		go jsonrpc.ServeConn(conn)
	}
}
