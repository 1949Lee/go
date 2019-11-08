package handler

import "leeBlogCli/test/fileServer"

type ParamNewArticle struct {
	Text string
	File fileServer.File
}
