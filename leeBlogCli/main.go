package main

import (
	"flag"
	"fmt"
	"leeBlogCli/concurrent"
	"leeBlogCli/config"
	"net/http"
	_ "net/http/pprof"
)

var a = 1

func main() {
	env := flag.String("env", "dev", "leeBlogCli's running environment")
	flag.Parse()
	config.ENV = *env
	fmt.Println(config.ENV)
	config.Init()
	lee := concurrent.Lee{}
	lee.Run()
	defer lee.Close()
	http.HandleFunc(config.WebsocketParserPath, lee.API.WebSocketReadMarkdownText)
	http.HandleFunc(config.WebsocketCheckLoginPath, lee.API.WebSocketCheckLogin)
	http.HandleFunc(config.ConfirmLogin, lee.API.APIInterceptor(lee.API.ConfirmLogin))
	http.HandleFunc(config.NewFile, lee.API.APIInterceptor(lee.API.ReceivingFile))
	http.HandleFunc(config.DeleteFile, lee.API.APIInterceptor(lee.API.DeleteFile))
	http.HandleFunc(config.NewArticleID, lee.API.APIInterceptor(lee.API.NewArticleID))
	http.HandleFunc(config.GetArticleWithEditingInfo, lee.API.APIInterceptor(lee.API.GetArticleWithEditingInfo))
	http.HandleFunc(config.SaveArticle, lee.API.APIInterceptor(lee.API.SaveArticle))
	http.HandleFunc(config.ArticleList, lee.API.APIInterceptor(lee.API.ArticleList))
	http.HandleFunc(config.ShowArticle, lee.API.APIInterceptor(lee.API.ShowArticle))
	http.HandleFunc(config.TagsGroupByCategory, lee.API.APIInterceptor(lee.API.GetTagsGroupByCategory))
	http.HandleFunc(config.NewTag, lee.API.APIInterceptor(lee.API.NewTag))
	http.HandleFunc(config.DeleteTag, lee.API.APIInterceptor(lee.API.DeleteTag))
	http.HandleFunc(config.NewCategory, lee.API.APIInterceptor(lee.API.NewCategory))
	http.HandleFunc(config.DeleteCategory, lee.API.APIInterceptor(lee.API.DeleteCategory))
	http.HandleFunc(config.FileResource, lee.API.ResourceInterceptor(lee.API.FileResource))
	fmt.Printf("server start with http://%s:%s\n", config.Self_URL, config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
