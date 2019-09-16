package engine

import (
	"goLearning/learn/helloworld/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, s := range seeds {
		requests = append(requests, s)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetcher: Start fetching url item %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: Error fetching url (%s) with result: %v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
