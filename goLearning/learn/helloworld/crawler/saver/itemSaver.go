package saver

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Got item #%d %+v", itemCount, item)
			//save(item)

		}
	}()

	return out
}

func save(item interface{}) {
	es, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := es.Index().Index("data_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
