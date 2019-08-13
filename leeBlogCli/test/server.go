package main

import (
	"fmt"
	"leeBlogCli/test/handler"
	"net/http"
)

func main() {
	port := "1314"
	http.HandleFunc("/", handler.ReadMarkdownText)
	fmt.Printf("server start with http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	//if file, err := os.Open("./test/test.md"); err != nil {
	//	panic(err)
	//} else {
	//	scanner := bufio.NewScanner(file)
	//	for scanner.Scan() {
	//		fmt.Println(scanner.Text())
	//	}
	//}
}
