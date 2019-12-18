package handler

import (
	"encoding/json"
	"leeBlogCli/definition"
	"log"
	"net/http"
)

func (api *API) NewArticleID(writer http.ResponseWriter, r *http.Request) {
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
	//for _, v := range r.MultipartForm.File {
	//	for i := 0; i < len(v); i++ {
	//		file, err := v[i].Open()
	//		if err != nil {
	//			log.Printf("ReceivingFile Handler when FormFile Error:%v", err)
	//			result.Code = 1
	//			result.Data = "上传文件打开失败"
	//		}
	//		buffer := bufio.NewReader(file)
	//		if err := os.MkdirAll(GetFilePath(param.ArticleID), os.ModePerm); err != nil {
	//			log.Printf("ReceivingFile Handler when os.MkdirAll Error:%v", err)
	//			result.Code = 1
	//			result.Data = "服务器保存文件失败"
	//		}
	//		f, err := os.Create(getFileName(param.ArticleID, v[i].Filename))
	//		if err != nil {
	//			log.Printf("ReceivingFile Handler when os.Create Error:%v", err)
	//			result.Code = 1
	//			result.Data = "服务器保存文件失败"
	//		}
	//		_, err = buffer.WriteTo(f)
	//		if err != nil {
	//			log.Printf("ReceivingFile Handler when buffer.WriteTo Error:%v", err)
	//			result.Code = 1
	//			result.Data = "服务器保存文件失败"
	//		}
	//		_ = file.Close()
	//	}
	//}
	var b []byte
	if b, err = json.Marshal(result); err != nil {
		log.Printf("ReceivingFile Handler when json.Marshal Error:%v", err)
		result.Code = 1
		result.Data = "服务器保存文件失败"
	}

	writer.Write([]byte(b))
}
