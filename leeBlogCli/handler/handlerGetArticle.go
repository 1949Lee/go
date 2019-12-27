package handler

import (
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/definition"
	"net/http"
)

func (api *API) GetArticleWithEditingInfo(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.FileOptions{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 2, Type: 4, Data: "参数获取失败"})
	}
	if param.ArticleID == "" {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章ID缺失"})
	}

	if param.FileName == "" {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文件名缺失"})
	}
	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}
	//api.Server.dasdasda
	_, _ = writer.Send(result)
}
