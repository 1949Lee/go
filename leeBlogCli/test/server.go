package main

import (
	"bufio"
	"fmt"
	"leeBlogCli/test/handler"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", handler.ReadMarkdownText)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	if file, err := os.Open("./test/test.md"); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
