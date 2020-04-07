package handler

import (
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/definition"
	"net/http"
)

// 获取登录的key
func (api *API) GetLoginKey(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	param := definition.ConfirmLoginParam{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "参数获取失败"})
		return
	}

	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}

	//list := api.Server.ArticleList(&param)
	//result.Data = list

	_, _ = writer.Send(result)
}

// 确认登录
func (api *API) ConfirmLogin(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	param := definition.ConfirmLoginParam{}
	err := json.Unmarshal(body, &param)
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 5, Data: "参数获取失败"})
		return
	}

	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}

	//list := api.Server.ArticleList(&param)
	//result.Data = list

	_, _ = writer.Send(result)
}
