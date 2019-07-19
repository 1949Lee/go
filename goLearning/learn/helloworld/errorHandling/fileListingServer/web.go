package main

import (
	"goLearning/learn/helloworld/errorHandling/fileListingServer/fileListing"
	"log"
	"net/http"
	"os"
)

type UserError interface {
	error
	Message() string
}

type appHandler func(http.ResponseWriter, *http.Request) error

func errorWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// 服务内部出错的最终拦截
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error handling request: %s", err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()

		err := handler(writer, request)

		// 如果路由对应的api返回已经处理过的错误
		if r, ok := err.(UserError); ok {
			http.Error(writer, r.Message(), http.StatusInternalServerError)
			return
		}

		// 发生一些可以预料，但是路由对应的api没有处理的错误
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			status := http.StatusOK
			switch {
			case os.IsNotExist(err):
				status = http.StatusNotFound
			case os.IsPermission(err):
				status = http.StatusForbidden
			default:
				status = http.StatusInternalServerError
			}
			text := http.StatusText(status)
			http.Error(writer, text, status)
			return
		}
	}
}

func main() {

	http.HandleFunc("/", errorWrapper(fileListing.FileList))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
