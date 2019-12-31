package handler

import (
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

// 根据传入的文章id和文件名得到文章对应的存放文件的路径
func GetDraftFilePath(articleID int) string {
	id := strconv.Itoa(articleID)
	builder := strings.Builder{}
	builder.WriteString(config.DraftPath)
	builder.WriteString(id)
	builder.WriteString(id)
	builder.WriteString(".draft")
	return builder.String()
}

func (api *API) GetArticleWithEditingInfo(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.GetArticleParam{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 2, Type: 4, Data: "参数获取失败"})
		return
	}
	if param.ArticleID == 0 {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章ID缺失"})
		return
	}

	if param.Type == 0 {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 4, Data: "参数文章状态缺失"})
		return
	}
	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}
	if err := os.MkdirAll(utils.GetDraftPath(int(param.ArticleID)), os.ModePerm); err != nil {
		log.Printf("GetArticleWithEditingInfo API 打开草稿所在目录时报错 Error:%v", err)
		result.Code = 1
		result.Data = "获取上次编辑结果失败"
		_, _ = writer.Send(result)
		return
	}
	resData := definition.EditingArticleInfo{}

	// 无论是未发布还是发布的文章，都首先获取草稿。
	filePath := GetDraftFilePath(int(param.ArticleID))
	draft := make([]byte, 0)
	_, err = os.Stat(filePath)
	if err == nil {
		draft, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("GetArticleWithEditingInfo API 打开草稿文件时报错 Error:%v", err)
			result.Code = 1
			result.Data = "获取上次编辑结果失败"
			_, _ = writer.Send(result)
			return
		}
	}

	// 发布文章的处理：需要返回两部分数据，文章的信息和文章的markdown。文章信息来自数据库，文章的markdown优先从草稿获取，草稿没有则从数据库获取。
	if param.Type == definition.GetArticleParamTypeEnum.PublicArticle {

		// 查询数据库获取文章信息
		article := api.Server.GetArticleHeader(&param)
		resData.Markdown = article.Content
		//resData.ArticleHeader.ID = article.ID
		//resData.ArticleHeader.CreateTime = article.CreateTime
		//resData.ArticleHeader.Title = article.Title
		//resData.ArticleHeader.Category = definition.Category{
		//    ID:article.CategoryID,
		//}
		if article.ID == 0 {
			result.Data = nil
		} else {
			result.Data = article
		}

		// 存在草稿，则使用草稿作为文章的markdown。否则需要到查询数据库的数据用做markdown。
		if len(draft) > 0 {

		}
	} else if param.Type == definition.GetArticleParamTypeEnum.DraftArticle { // 未发布的文章的处理：

	}
	//result.Data = resData
	_, _ = writer.Send(result)
}

func (api *API) ArticleList(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	param := definition.ArticleListParam{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "参数获取失败"})
		return
	}

	if param.PageSize == 0 {
		param.PageSize = config.IndexArticleListPageSize
		//_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "参数文章类型缺失"})
		//return
	}

	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}
	list := definition.ArticleListResult{}

	result.Data = list

	ok := api.Server.ArticleList(&param)
	if !ok {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "发布失败"})
		return
	}

	_, _ = writer.Send(result)
}