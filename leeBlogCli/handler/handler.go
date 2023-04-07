package handler

import (
	"github.com/gorilla/websocket"
	"leeBlogCli/definition"
	"leeBlogCli/server"
	"net/http"
)

type API struct {
	Server    *server.Blog
	LoginConn *websocket.Conn
}

//// 检查登录，检验登录通过返回true，未通过返回false。code码返回 0，登录检验通过；返回 1，参数错误；返回 2，用户未登录；返回 3，token无效。
//func (api *API) CheckLogin(body []byte) (bool,int) {
//    param := map[string]interface{}{}
//    err := json.Unmarshal(body, &param)
//    if err  != nil {
//        // 参数错误
//        return false, 1
//    }
//    if param["leeKey"] == nil || param["leeToken"] == nil {
//        // 未传登录凭证
//        return false, 1
//    } else if param["leeKey"] == "" || param["leeToken"] == "" {
//        // 未传登录凭证
//        return false, 1
//    }
//    leeKey, ok := param["leeKey"].(string)
//    if !ok {
//        // 参数错误
//        return false, 1
//    }
//    leeToken, err := param["leeToken"].(string)
//    if !ok {
//        // 参数错误
//        return false, 1
//    }
//    token, ok := api.Server.LeeTokenMap[leeKey]
//    if !ok {
//        // 用户未登录
//       return false, 2
//    }
//    if token != leeToken {
//        // token无效
//        return false, 3
//    }
//    return true, 0
//}

// 检查登录，检验登录通过返回true，未通过返回false。
func (api *API) CheckLogin(r *http.Request) bool {
	leeKey := r.Header.Get("leeKey")
	leeToken := r.Header.Get("leeToken")
	if leeKey == "" || leeToken == "" {
		return false
	}

	token, ok := api.Server.LeeTokenMap[leeKey]
	if !ok {
		return false
	}

	if token != leeToken {
		return false
	}
	return true
}

// 404
func (api *API) NotFound(writer *APIResponseWriter, r *http.Request) {
	result := definition.APIResult{
		Code:    404,
		Data:    "",
		Message: "不合理请求",
	}
	_, _ = writer.SendOtherStatus(result, http.StatusNotFound)
}
