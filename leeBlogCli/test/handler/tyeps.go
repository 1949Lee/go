package handler

import (
	"leeBlogCli/test/fileServer"
)

type ParamNewArticle struct {
	Type  uint8
	Text  string
	Files []fileServer.File
}
