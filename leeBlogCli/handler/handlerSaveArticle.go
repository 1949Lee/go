package handler

import (
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/definition"
	"net/http"
)

// 新增或更新文章
func (api *API) SaveArticle(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.SaveArticleInfo{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "参数获取失败"})
		return
	}

	if param.Type == 0 {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "参数文章类型缺失"})
		return
	}
	if param.Content == "" {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章markdown字符串缺失"})
		return
	}
	if param.Text == "" {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章text缺失"})
		return
	}
	if param.ArticleHeader.ID == 0 {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章ID缺失"})
		return
	}
	if param.ArticleHeader.Category == (definition.Category{}) {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章所属分类缺失"})
		return
	}
	if len(param.ArticleHeader.Tags) == 0 {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章所属标签缺失"})
		return
	}
	if param.ArticleHeader.Title == "" {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章标题缺失"})
		return
	}
	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}

	ok := api.Server.SaveArticle(&param)
	if !ok {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "发布失败"})
		return
	}

	_, _ = writer.Send(result)
}
