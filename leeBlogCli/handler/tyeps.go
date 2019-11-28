package handler

import "leeBlogCli/concurrent"

type ParamNewArticle struct {
	Type  uint8
	Text  string
	Files []concurrent.File
}
