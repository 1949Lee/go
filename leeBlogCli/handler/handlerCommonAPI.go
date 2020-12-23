package handler

import (
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/definition"
	"log"
	"net/http"
)

// 生成新的文章ID的API
func (api *API) NewArticleID(writer *APIResponseWriter, r *http.Request) {
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	result.Data = api.Server.GenerateArticleID()
	_, _ = writer.Send(result)
}

func (api *API) GetTagsGroupByCategory(writer *APIResponseWriter, r *http.Request) {
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	result.Data = api.Server.GetTagsGroupByCategory()
	_, _ = writer.Send(result)
}

// 添加新标签的API
func (api *API) NewTag(writer *APIResponseWriter, r *http.Request) {
	param := definition.Tag{
		ID: -1,
	}
	rData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%v", err)
	}
	err = json.Unmarshal(rData, &param)
	if err != nil {
		log.Printf("%v", err)
	}
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	result.Data = api.Server.NewTag(param)
	_, _ = writer.Send(result)
}

// 删除标签的API
func (api *API) DeleteTag(writer *APIResponseWriter, r *http.Request) {
	param := definition.Tag{
		ID: -1,
	}
	rData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%v", err)
	}
	err = json.Unmarshal(rData, &param)
	if err != nil {
		log.Printf("%v", err)
	}
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	ok := api.Server.DeleteTag(param)
	if !ok {
		result.Code = 1
		result.Message = "删除失败"
		result.Data = nil
	}
	_, _ = writer.Send(result)
}

// 添加新分类的API
func (api *API) NewCategory(writer *APIResponseWriter, r *http.Request) {
	param := definition.Category{
		ID: -1,
	}
	rData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%v", err)
	}
	err = json.Unmarshal(rData, &param)
	if err != nil {
		log.Printf("%v", err)
	}
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	result.Data = api.Server.NewCategory(param)
	_, _ = writer.Send(result)
}

// 删除分类的API
func (api *API) DeleteCategory(writer *APIResponseWriter, r *http.Request) {

	param := definition.Category{
		ID: -1,
	}
	rData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%v", err)
	}
	err = json.Unmarshal(rData, &param)
	if err != nil {
		log.Printf("%v", err)
	}
	result := definition.APIResult{
		Code: 0,
		Data: nil,
	}
	defer r.Body.Close()

	result.Data = api.Server.DeleteCategory(param)
	_, _ = writer.Send(result)
}
