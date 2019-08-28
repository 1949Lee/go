package parser

import (
	"goLearning/learn/helloworld/crawler/engine"
	"regexp"
)

const cityListRegexp string = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func CityListParser(body []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRegexp)
	matches := re.FindAllSubmatch(body, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: engine.NilParserFunc})
	}
	return result
}
