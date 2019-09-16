/**
爬取珍爱网的爬虫
*/

package main

import (
	"bufio"
	"goLearning/learn/helloworld/crawler/engine"
	"goLearning/learn/helloworld/crawler/scheduler"
	"goLearning/learn/helloworld/crawler/zhenai/parser"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	//engine.SimpleEngine{}.Run(engine.Request{Url: url, ParserFunc: parser.CityListParser})
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 1000,
	}
	e.Run(engine.Request{Url: url, ParserFunc: parser.CityListParser})

}

/**
如果url返回的不是utf8编码的话，需要进行转换。这里距离说明gbk如何转化为utf8
1. 安装转换用到的工具包：
	gopm get -g -v golang.org/x/text
	gopm get -g -v golang.org/x/net/html
2. 代码：
	e := determineEncoding(resp.Body)
	utf8reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8reader)
*/
//获取读取的网页内容的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
