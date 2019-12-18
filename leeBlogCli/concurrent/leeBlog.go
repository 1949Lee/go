package concurrent

import "leeBlogCli/dao"

type Blog struct {
	dao *dao.DBServer
}

func (b *Blog) Run() {
	db := dao.DBServer{}
	db.Open()
	b.dao = &db
}

func (b *Blog) Close() {
	b.dao.Close()
}
