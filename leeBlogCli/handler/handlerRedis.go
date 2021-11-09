package handler

import (
	"encoding/json"
	"io/ioutil"
	"leeBlogCli/definition"
	"net/http"
)

// GetRedisValueByKey 根据key获取redis的值
func (api *API) GetRedisValueByKey(writer *APIResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	param := definition.GetRedisValueByKeyParam{}
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

	list := api.Server.GetRedisValueByKey(&param)
	result.Data = list

	_, _ = writer.Send(result)
}
func (api *API) GetImageTypeNeedConvert() string {
	return api.Server.GetImageTypeNeedConvert()
}

// InitRedis 初始化redis的数据
func (api *API) InitRedis(writer *APIResponseWriter, r *http.Request) {
	ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	result := definition.ResponseResult{
		Type: 4,
		Code: 0,
		Data: "成功",
	}

	err := api.Server.InitRedis()
	if err != nil {
		_, _ = writer.Send(definition.ResponseResult{Code: 1, Type: 1, Data: "初始化redis数据失败"})
		return
	}
	result.Data = ""

	_, _ = writer.Send(result)
}
