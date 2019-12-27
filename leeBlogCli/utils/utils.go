package utils

import (
	"leeBlogCli/config"
	"strconv"
	"strings"
)

// 根据传入的文章id和文件名得到文章对应的存放文件的路径
func GetFilePath(articleID int) string {
	builder := strings.Builder{}
	builder.WriteString(config.FilePath)
	builder.WriteString(strconv.Itoa(articleID))
	//builder.WriteString("/")
	//builder.WriteString(fileName)
	return builder.String()
}
