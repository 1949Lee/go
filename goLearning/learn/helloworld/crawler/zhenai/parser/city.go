package parser

import (
	"goLearning/learn/helloworld/crawler/engine"
	"regexp"
)

const cityRegexp string = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func CityParser(body []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRegexp)
	matches := re.FindAllSubmatch(body, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		name := string(m[2])
		//func(m [][]byte) {
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: func(bytes []byte) engine.ParserResult {
			return ProfileParser(bytes, name)
		}})
		//}(m)
	}
	return result
}
