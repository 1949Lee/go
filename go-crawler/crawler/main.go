package main

import (
	"go-crawler/crawler/config"
	"go-crawler/crawler/engine"
	"go-crawler/crawler/persist"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
