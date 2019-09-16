package parser

import (
	"goLearning/learn/helloworld/crawler/engine"
	"goLearning/learn/helloworld/crawler/model"
	"regexp"
)

const profileRegexp string = `<div class="id"[^>]*>ID：([^<]+)</div>`

//func ProfileParser(body []byte, params map[string]string) engine.ParserResult {
func ProfileParser(body []byte, name string) engine.ParserResult {
	re := regexp.MustCompile(profileRegexp)
	str := re.FindSubmatch(body)
	result := engine.ParserResult{}
	user := model.Profile{Name: name}
	if len(str) > 1 {
		user.ID = string(str[1])
	}
	result.Items = []interface{}{user}
	result.Requests = []engine.Request{}
	return result
}
