package concurrent

import (
	"leeBlogCli/dao"
	"leeBlogCli/handler"
	"leeBlogCli/server"
)

type Lee struct {
	API *handler.API
}

func (b *Lee) Run() {
	daoServer := dao.DBServer{}
	daoServer.Open()
	blog := server.Blog{
		Dao:         &daoServer,
		LeeTokenMap: map[string]string{},
	}
	api := handler.API{
		Server: &blog,
	}
	b.API = &api
}

func (b *Lee) Close() {
	b.API.Server.Dao.Close()
}
