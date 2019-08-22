package main

import (
	"fmt"
	"leeBlogCli/test/handler"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	port := "1314"
	http.HandleFunc("/", handler.ReadMarkdownText)
	fmt.Printf("server start with http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
