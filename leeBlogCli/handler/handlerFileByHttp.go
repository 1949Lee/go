package handler

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"leeBlogCli/utils"
	"log"
	"net/http"
	"os"
	"strconv"
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

// 根据传入的文章id和文件名得到最终文件名
func getFileSourceName(articleID string, fileName string) string {
	builder := strings.Builder{}
	builder.WriteString(config.FileSourcePath)
	builder.WriteString(articleID)
	builder.WriteString("/")
	builder.WriteString(fileName)
	return builder.String()
}

// 接收上传的文件
func (api *API) ReceivingFile(writer *APIResponseWriter, r *http.Request) {
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
	fileres := definition.FileResponse{}
	for _, v := range r.MultipartForm.File {
		for i := 0; i < len(v); i++ {
			fileResItem := definition.FileResponseItem{
				FileName: v[i].Filename,
				URL:      config.APIFullURL + config.FileResource + getFileSourceName(param.ArticleID, v[i].Filename),
			}
			file, err := v[i].Open()
			if err != nil {
				log.Printf("ReceivingFile Handler when FormFile Error:%v", err)
				result.Code = 1
				result.Data = "上传文件打开失败"
				fileResItem.URL = ""
			}
			buffer := bufio.NewReader(file)
			articleID, atoIErr := strconv.Atoi(param.ArticleID)
			if atoIErr == nil {
				if err := os.MkdirAll(utils.GetFilePath(articleID), os.ModePerm); err != nil {
					log.Printf("ReceivingFile Handler when os.MkdirAll Error:%v", err)
					result.Code = 1
					result.Data = "服务器保存文件失败"
					fileResItem.URL = ""
				}
			} else {
				log.Printf("上传文件接口获取参数错误，文章ID值不合法")
				result.Code = 1
				result.Data = "参数文章ID值不合法"
				fileResItem.URL = ""
			}
			f, err := os.Create(getFileName(param.ArticleID, v[i].Filename))
			if err != nil {
				log.Printf("ReceivingFile Handler when os.Create Error:%v", err)
				result.Code = 1
				result.Data = "服务器保存文件失败"
				fileResItem.URL = ""
			}
			_, err = buffer.WriteTo(f)
			if err != nil {
				log.Printf("ReceivingFile Handler when buffer.WriteTo Error:%v", err)
				result.Code = 1
				result.Data = "服务器保存文件失败"
				fileResItem.URL = ""
			}
			_ = file.Close()
			// TODO 执行服务器命令来转换为渐进式图片，质量参数可变做成配置项：convert 1.jpg -quality 80  -interlace plane 2.jpg
			if fileResItem.URL != "" {
				fileres = append(fileres, fileResItem)
			}
		}
	}

	if len(fileres) > 0 {
		result.Data = fileres
	}

	_, _ = writer.Send(result)
}

// 删除上传的文件
func (api *API) DeleteFile(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.FileOptions{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 2, Type: 4, Data: "参数获取失败"})
		return
	}
	if param.ArticleID == "" {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章ID缺失"})
		return
	}

	if param.FileName == "" {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文件名缺失"})
		return
	}
	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "删除成功",
	}
	err = os.Remove(getFileName(param.ArticleID, param.FileName))
	if err != nil {
		log.Printf("删除文件失败，error：%v ", err)
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "删除失败"})
		return
	}
	articleID, atoIErr := strconv.Atoi(param.ArticleID)
	if atoIErr != nil {
		log.Printf("删除文章最后一个文件后，删除文件夹时转换文章ID失败，error：%v ", atoIErr)
		result.Code = 1
		result.Data = "删除失败"
		_, _ = writer.Send(result)
		return
	} else {
		path := utils.GetFilePath(articleID)
		infos, err := ioutil.ReadDir(path)
		if err != nil {
			log.Printf("删除文章最后一个文件后，删除文件夹时读取文件夹失败，error：%v ", err)
			result.Code = 1
			result.Data = "删除失败"
			_, _ = writer.Send(result)
			return
		}
		if len(infos) == 0 {
			err := os.RemoveAll(path)
			if err != nil {
				log.Printf("删除文章最后一个文件后，删除文件夹失败，error：%v ", err)
				result.Code = 1
				result.Data = "删除失败"
				_, _ = writer.Send(result)
				return
			}
		}
	}

	_, _ = writer.Send(result)
}

// 获取文章的上传的文件列表
func (api *API) GetArticleFileList(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.ArticleIDParam{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 2, Type: 4, Data: "参数获取失败"})
		return
	}
	if param.ArticleID == 0 {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章ID缺失"})
		return
	}
	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}

	resData := definition.ArticleFileList{}

	// 获取文章的已上传文件列表。
	infos, err := ioutil.ReadDir(utils.GetFilePath(int(param.ArticleID)))
	resData.List = make(map[string]definition.ArticleFileListItem)
	for _, f := range infos {
		if !f.IsDir() {
			resData.List[f.Name()] = definition.ArticleFileListItem{
				Name: f.Name(),
				Size: int(f.Size()),
				Url:  config.APIFullURL + config.FileResource + getFileSourceName(strconv.Itoa(int(param.ArticleID)), f.Name()),
			}
		}
	}

	result.Data = resData
	_, _ = writer.Send(result)
}
