package handler

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"log"
	"net/http"
	"os"
	"strings"
)

// 根据传入的文章id和文件名得到最终文件名
func getFileName(articleID string, fileName string) string {
	builder := strings.Builder{}
	builder.WriteString(config.FilePath)
	builder.WriteString(articleID)
	builder.WriteString("/")
	builder.WriteString(fileName)
	return builder.String()
}

// 根据传入的文章id和文件名得到最终文件路径
func getFilePath(articleID string) string {
	builder := strings.Builder{}
	builder.WriteString(config.FilePath)
	builder.WriteString(articleID)
	//builder.WriteString("/")
	//builder.WriteString(fileName)
	return builder.String()
}

// 接收上传的文件
func ReceivingFile(writer http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	/**
	  此部分代码可以提取出来。*/
	//设置跨域的相应头CORS，CORS参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS
	writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	// 请求方发送请求时当请求为options时，直接返回200。
	if r.Method == "OPTIONS" {
		writer.WriteHeader(200)
		return
	}
	//下面是接口真正的处理
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "接收成功",
	}
	err := r.ParseMultipartForm(64 << 10)
	defer r.Body.Close()
	if err != nil {
		log.Printf("ReceivingFile Handler when ParseMultipartForm Error:%v", err)
		result.Code = 1
		result.Data = "上传失败"
	}
	param := definition.FileOptions{}
	param.ArticleID = r.FormValue("articleID")
	if param.ArticleID == "" {
		log.Printf("上传文件接口获取参数错误，缺少文章ID")
		result.Code = 1
		result.Data = "参数文章ID缺失"
	}
	for _, v := range r.MultipartForm.File {
		for i := 0; i < len(v); i++ {
			file, err := v[i].Open()
			if err != nil {
				log.Printf("ReceivingFile Handler when FormFile Error:%v", err)
				result.Code = 1
				result.Data = "上传文件打开失败"
			}
			buffer := bufio.NewReader(file)
			if err := os.MkdirAll(getFilePath(param.ArticleID), os.ModePerm); err != nil {
				log.Printf("ReceivingFile Handler when os.MkdirAll Error:%v", err)
				result.Code = 1
				result.Data = "服务器保存文件失败"
			}
			f, err := os.Create(getFileName(param.ArticleID, v[i].Filename))
			if err != nil {
				log.Printf("ReceivingFile Handler when os.Create Error:%v", err)
				result.Code = 1
				result.Data = "服务器保存文件失败"
			}
			_, err = buffer.WriteTo(f)
			if err != nil {
				log.Printf("ReceivingFile Handler when buffer.WriteTo Error:%v", err)
				result.Code = 1
				result.Data = "服务器保存文件失败"
			}
			_ = file.Close()
		}
	}
	var b []byte
	if b, err = json.Marshal(result); err != nil {
		log.Printf("ReceivingFile Handler when json.Marshal Error:%v", err)
		result.Code = 1
		result.Data = "服务器保存文件失败"
	}

	writer.Write([]byte(b))
}

// 删除上传的文件
func DeleteFile(writer http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	/**
	  此部分代码可以提取出来。*/
	//设置跨域的相应头CORS，CORS参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS
	writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	// 请求方发送请求时当请求为options时，直接返回200。
	if r.Method == "OPTIONS" {
		writer.WriteHeader(200)
		return
	}
	//下面是接口真正的处理
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.FileOptions{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		b, _ := json.Marshal(definition.ResponseResult{Code: 2, Type: 4, Data: "参数获取失败"})
		writer.Write(b)
	}
	if param.ArticleID == "" {
		b, _ := json.Marshal(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章ID缺失"})
		writer.Write(b)
	}

	if param.FileName == "" {
		b, _ := json.Marshal(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文件名缺失"})
		writer.Write(b)
	}
	err = os.Remove(getFileName(param.ArticleID, param.FileName))
	if err != nil {
		b, _ := json.Marshal(definition.ResponseResult{Code: 1, Type: 4, Data: "删除失败"})
		log.Printf("删除文件失败，error：%v ", err)
		writer.Write(b)
	}

	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "删除成功",
	}
	var b []byte
	if b, err = json.Marshal(result); err != nil {
		log.Printf("ReceivingFile Handler when json.Marshal Error:%v", err)
	}

	writer.Write([]byte(b))
}
