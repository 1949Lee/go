package server

import "leeBlogCli/dao"

type Blog struct {
	Dao *dao.DBServer

	// leeToken的数组（单机版不用redis）
	LeeTokenMap map[string]string
}
