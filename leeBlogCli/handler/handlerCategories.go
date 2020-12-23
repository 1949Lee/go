package handler

import (
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

	result.Data = api.Server.GetTagsWithArticleID()
	_, _ = writer.Send(result)
}
