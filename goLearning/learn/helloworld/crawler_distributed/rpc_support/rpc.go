package rpc_support

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 开启一个新的rpc服务
func ServeRPC(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":"+host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}

		// 开一个协程去执行，否则会祖册下一个调用rpc的人
		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", ":"+host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
