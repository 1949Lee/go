package handler

import (
	"io/ioutil"
	"leeBlogCli/config"
	"log"
	"net/http"
	"os"
)

func (api *API) FileResource(writer *APIResponseWriter, r *http.Request) {
	requestUrl := r.URL.Path
	filePath := requestUrl[len(config.FileResource):]
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Println("static resource:", err)
		writer.writer.WriteHeader(404)
	} else {
		bs, _ := ioutil.ReadAll(file)

		_, _ = writer.Write(bs)
	}
}
