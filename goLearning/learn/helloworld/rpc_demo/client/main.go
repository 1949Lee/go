package main

import (
	"fmt"
	"goLearning/learn/helloworld/rpc_demo"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1949")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpc_demo.Args{A: 3, B: 0}, &result)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(result)
	}
}
