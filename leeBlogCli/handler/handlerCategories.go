package handler

import (
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"net/http"
)

// 获取所有分类和分类下的所有标签，每个标签附带含有标签的文章ID数组
func (api *API) GetTagsWithArticleID(writer *APIResponseWriter, r *http.Request) {
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	list := api.Server.GetTagsWithArticleID()
	result.Data = list
	_, _ = writer.Send(result)
}

// 根据文章ID获取文章列表
func (api *API) GetArticleListByID(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.ArticleListByIDParam{}
	param.CategoryID = -1
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "参数获取失败"})
		return
	}

	// 默认也大小
	if param.PageSize == 0 {
		param.PageSize = config.IndexArticleListPageSize
	}

	if param.CategoryID == -1 {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "无法搜索全部分类，请确定选择了文章分类"})
		return
	}

	// 默认页码为1
	if param.PageIndex == 0 {
		param.PageIndex = 1
	}

	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}

	list := api.Server.GetArticleListByID(&param)
	result.Data = list

	_, _ = writer.Send(result)
}
