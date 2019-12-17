package utils

import (
	"leeBlogCli/config"
	"strconv"
	"strings"
)

// 根据传入的文章id和文件名得到最终文件路径
func GetFilePath(articleID int) string {
	builder := strings.Builder{}
	builder.WriteString(config.FilePath)
	builder.WriteString(strconv.Itoa(articleID))
	//builder.WriteString("/")
	//builder.WriteString(fileName)
	return builder.String()
}
