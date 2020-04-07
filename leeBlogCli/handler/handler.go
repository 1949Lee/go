package handler

import (
	"github.com/gorilla/websocket"
	"leeBlogCli/server"
)

type API struct {
	Server    *server.Blog
	LoginConn *websocket.Conn
}
