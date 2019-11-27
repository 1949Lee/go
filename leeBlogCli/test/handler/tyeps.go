package handler

import (
	"leeBlogCli/test/concurrent"
)

type ParamNewArticle struct {
	Type  uint8
	Text  string
	Files []concurrent.File
}
