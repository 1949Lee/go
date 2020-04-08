package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"log"
	"net/http"
	"strings"
)

var loginUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		if strings.HasSuffix(r.Header.Get("Origin"), config.LegalOriginURL) {
			return true
		}
		return false
	},
}

// 扫描二维码登录轮训的WebSocket接口
func (api *API) WebSocketCheckLogin(writer http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(writer, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	api.LoginConn = conn
	defer func() {
		_ = conn.Close()
		api.LoginConn = nil
	}()
	api.websocketLoginLoop(conn)
}

// 扫描二维码登录轮训的WebSocket接口的主循环
func (api *API) websocketLoginLoop(conn *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("websocketLoginLoop error %s", err)
			//if err = conn.WriteJSON(conn.WriteJSON(ResponseResult{Code: 1})); err != nil {
			//	log.Printf("write err:%v", err)
			//	//return
			//}
			api.websocketLoginLoop(conn)
		}
	}()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			_, ok := err.(*websocket.CloseError)
			if ok {
				break
			}
			log.Printf("receive err:%v", err)
			continue
		}
		if messageType == websocket.TextMessage {
			// 收到网站的任意消息，则返回管理员email
			if string(p) != "" {
				result := definition.APIResult{
					Code: 0,
					Data: nil,
				}

				// 获取登录的key(管理员email)
				result.Data = api.Server.GetLoginKey()
				if err := conn.WriteJSON(result); err != nil {
					log.Printf("write err:%v", err)
				}
			}
			//obj.Type
		}
	}
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

	result := definition.APIResult{
		Code: 0,
		Data: "成功",
	}

	data := api.Server.ConfirmLoginInfo(&param)
	if data.LeeToken == "-1" {
		result.Data = nil
		result.Message = "登录失败"
		result.Code = 1
	} else {
		result.Data = data
		api.Server.LeeTokenMap[param.Email] = data.LeeToken
	}

	// 通知WebSocketCheckLogin接口登录失败/成功
	if api.LoginConn != nil {
		if err := api.LoginConn.WriteJSON(result); err != nil {
			log.Printf("write err:%v", err)
		}
	}

	_, _ = writer.Send(result)
}
