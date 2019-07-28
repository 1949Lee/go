package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic("panic error")
}

func errUserErr(writer http.ResponseWriter, request *http.Request) error {
	return userError("user error")
}

func noErr(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprint(writer, "no error")
	return nil
}

func noNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func noPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func unknowErr(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrInvalid
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{h: errPanic, code: 500, message: "Internal Server Error"},
	{h: errUserErr, code: 500, message: "user error"},
	{h: noErr, code: 200, message: "no error"},
	{h: noNotFound, code: 404, message: "Not Found"},
	{h: noPermission, code: 403, message: "Forbidden"},
	{h: unknowErr, code: 500, message: "Internal Server Error"},
}

// 模拟测试webApi，用到了httptest中NewRecorder（response）和NewRequest（设置请求信息）
func TestErrorWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet, "https://www.jiaxuanlee.com", nil)
		f(res, req)
		b, _ := ioutil.ReadAll(res.Body)
		body := strings.Trim(string(b), "\n")
		if res.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s), but got (%d, %s)", tt.code, tt.message, res.Code, body)
		}

	}
}

// 用http启一个真正的server去测试webApi
func TestAPIFileListInServer(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		//起一个真正的服务
		server := httptest.NewServer(http.HandlerFunc(f))

		//发送http请求
		res, _ := http.Get(server.URL)
		b, _ := ioutil.ReadAll(res.Body)
		body := strings.Trim(string(b), "\n")
		if res.StatusCode != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s), but got (%d, %s)", tt.code, tt.message, res.StatusCode, body)
		}

	}
}
