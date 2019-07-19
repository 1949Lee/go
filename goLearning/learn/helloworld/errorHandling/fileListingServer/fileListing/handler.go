package fileListing

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func FileList(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("File list path is illegal!")
	}

	// 获取接口的path
	filePath := request.URL.Path[len(prefix):]
	//解析接口的path并寻找对应文件
	file, err := os.Open("errorHandling/" + filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
