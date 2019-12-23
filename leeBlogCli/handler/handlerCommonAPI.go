package handler

import (
	"encoding/json"
	"leeBlogCli/definition"
	"log"
	"net/http"
)

// 生成新的文章ID的API
func (api *API) NewArticleID(writer http.ResponseWriter, r *http.Request) {
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	result.Data = api.Server.GenerateArticleID()
	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("NewArticleID接口返回数据，转换json失败:%v", err)
		result.Code = 1
		result.Message = "获取新ID失败"
	}

	writer.Write(b)
}

// 生成新的文章ID的API
func (api *API) GetTagsGroupByCategory(writer http.ResponseWriter, r *http.Request) {
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	result.Data = api.Server.GetTagsGroupByCategory()
	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("GetTagsGroupByCategory接口返回数据，转换json失败:%v", err)
		result.Code = 1
		result.Message = "获取错误"
		result.Data = nil
	}

	writer.Write(b)
}
