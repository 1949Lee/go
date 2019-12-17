package concurrent

import "leeBlogCli/dao"

type Blog struct {
	dao *dao.DBServer
}

func (b *Blog) Run() {
	db := dao.DBServer{}
	db.Open()
}

func (b *Blog) Close() {
	defer b.dao.Close()
}
