package parser

import (
	"io/ioutil"
	"leeBlogCli"
	"os"
	"testing"
)

//性能测试，结果注意单位，ns表示纳秒
func BenchmarkMarkdownParse(t *testing.B) {
	file, err := os.Open("../语法格式.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	md, err := ioutil.ReadAll(file)
	//ans := 9
	t.ResetTimer()

	//性能测试循环测试的次数，由go test自动计算（t.N）
	for i := 0; i < t.N; i++ {
		tokenList, html := leeBlogCli.MarkdownParse(string(md))

		if html == "" || tokenList == nil {
			t.Error("parse error")
		}
		//l.
	}
	//if actual := l.p; actual != ans {
	//    t.Errorf("input string %s got %d; expected %d.", s, actual, ans)
	//}
}
