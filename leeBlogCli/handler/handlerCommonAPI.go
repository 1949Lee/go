package handler

import (
	"encoding/json"
	"io/ioutil"
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

// 添加新标签的API
func (api *API) NewTag(writer http.ResponseWriter, r *http.Request) {
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
	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("NewTag接口返回数据，转换json失败:%v", err)
		result.Code = 1
		result.Message = "添加失败"
		result.Data = nil
	}

	writer.Write(b)
}

// 删除标签的API
func (api *API) DeleteTag(writer http.ResponseWriter, r *http.Request) {
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
	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("DeleteTag接口返回数据，转换json失败:%v", err)
		result.Code = 1
		result.Message = "删除失败"
		result.Data = nil
	}

	writer.Write(b)
}

// 添加新分类的API
func (api *API) NewCategory(writer http.ResponseWriter, r *http.Request) {
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
	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("NewCategory接口返回数据，转换json失败:%v", err)
		result.Code = 1
		result.Message = "添加失败"
		result.Data = nil
	}

	writer.Write(b)
}

// 删除分类的API
func (api *API) DeleteCategory(writer http.ResponseWriter, r *http.Request) {

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
	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("DeleteCategory接口返回数据，转换json失败:%v", err)
		result.Code = 1
		result.Message = "删除失败"
		result.Data = nil
	}

	writer.Write(b)
}
