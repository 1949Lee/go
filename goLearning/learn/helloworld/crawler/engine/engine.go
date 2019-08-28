package engine

import (
	"goLearning/learn/helloworld/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, s := range seeds {
		requests = append(requests, s)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetcher: Start fetching url item %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: Error fetching url (%s) with result: %v", r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(body)

		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
