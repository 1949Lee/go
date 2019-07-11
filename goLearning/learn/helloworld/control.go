package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""

	// switch 后可以不跟表达式。每个case单独去写。switc语句的case可以不写break，go语言会自动帮你加
	switch {
	case score < 0:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	case score == 100:
		g = "S"
	default:
		panic(fmt.Sprintf("wrong score: %d", score)) // 终端程序的执行
	}
	return g
}

func main() {

	// 读取文件并输出文件内容
	//const filename = "abc.txt"
	const filename = "note.md"

	//写法一
	fmt.Println("*******写法一******")
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//写法二
	fmt.Println("******写法二******")
	//contents1的作用域是if-else语句
	if contents1, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents1)
	}

	//switch语句
	fmt.Println(grade(0), grade(59), grade(70), grade(91), grade(100))


	// 循环语句
}
