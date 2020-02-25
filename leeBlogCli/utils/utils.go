package utils

import (
	"leeBlogCli/config"
	"os"
	"strconv"
	"strings"
)

// 获取网站静态资源根路径
func GetRootDir() string {
	builder := strings.Builder{}
	homeDir, _ := os.UserHomeDir()
	builder.WriteString(homeDir)
	builder.WriteRune(os.PathSeparator)
	return builder.String()
}

// 根据传入的文章id和文件名得到文章对应的存放文件的路径
func GetFilePath(articleID int) string {
	builder := strings.Builder{}
	builder.WriteString(GetRootDir())
	builder.WriteString(config.FilePath)
	builder.WriteString(strconv.Itoa(articleID))
	//builder.WriteString("/")
	//builder.WriteString(fileName)
	return builder.String()
}

// 根据传入的文章id和文件名得到文章对应的存放文件的路径
func GetDraftPath(articleID int) string {
	builder := strings.Builder{}
	builder.WriteString(GetRootDir())
	builder.WriteString(config.DraftPath)
	builder.WriteString(strconv.Itoa(articleID))
	return builder.String()
}
