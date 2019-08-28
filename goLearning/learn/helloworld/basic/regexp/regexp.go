package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is lijiaxuan1001@gmail.com
My email is abc@gmail.com
My email is csda@gmail.com.cn
`

func main() {
	//re := regexp.MustCompile("lijiaxuan1001@gmail.com")
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[.a-zA-Z0-9]+)`)
	// 找出所有的email地址
	match := re.FindAllString(text, -1)
	fmt.Println(match)

	//邮箱的用户名和域名分别提取
	match2 := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match2)
	fmt.Println()
	fmt.Println("子匹配结果")
	for _, m := range match2 {
		fmt.Println(m)
	}
}
