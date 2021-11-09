package main

import (
	"flag"
	"fmt"
	"leeBlogCli/concurrent"
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"net/http"
	_ "net/http/pprof"
)

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
	http.HandleFunc(config.ConfirmLogin, lee.API.APIInterceptor(lee.API.ConfirmLogin, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.NewFile, lee.API.APIInterceptor(lee.API.ReceivingFile, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.DeleteFile, lee.API.APIInterceptor(lee.API.DeleteFile, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.NewArticleID, lee.API.APIInterceptor(lee.API.NewArticleID, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.GetArticleWithEditingInfo, lee.API.APIInterceptor(lee.API.GetArticleWithEditingInfo, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.SaveArticle, lee.API.APIInterceptor(lee.API.SaveArticle, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.ArticleList, lee.API.APIInterceptor(lee.API.ArticleList, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.GetArticleListByID, lee.API.APIInterceptor(lee.API.GetArticleListByID, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.ShowArticle, lee.API.APIInterceptor(lee.API.ShowArticle, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.TagsGroupByCategory, lee.API.APIInterceptor(lee.API.GetTagsGroupByCategory, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.TagsWithArticleID, lee.API.APIInterceptor(lee.API.GetTagsWithArticleID, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.NewTag, lee.API.APIInterceptor(lee.API.NewTag, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.DeleteTag, lee.API.APIInterceptor(lee.API.DeleteTag, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.NewCategory, lee.API.APIInterceptor(lee.API.NewCategory, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.DeleteCategory, lee.API.APIInterceptor(lee.API.DeleteCategory, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.GetRedisValueByKey, lee.API.APIInterceptor(lee.API.GetRedisValueByKey, definition.InterceptorOptions{CheckLogin: true}))
	http.HandleFunc(config.InitRedis, lee.API.APIInterceptor(lee.API.InitRedis, definition.InterceptorOptions{CheckLogin: false}))
	http.HandleFunc(config.FileResource, lee.API.ResourceInterceptor(lee.API.FileResource))
	fmt.Printf("server start with http://%s:%s\n", config.Self_URL, config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
