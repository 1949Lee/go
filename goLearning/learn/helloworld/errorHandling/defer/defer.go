package main

import (
	"bufio"
	"errors"
	"fmt"
	"goLearning/learn/helloworld/functional/fib"
	"os"
)

/**
defer语句会在函数运行完毕时运行：defer执行顺序为栈-先进后出，defer语句相关变量的计算会在定义的地方执行。
*/

/*
1. 注意defer1的顺序和中间数字的执行先后，这个顺序验证了defer调用是栈的方式，先进后出
2. 中间defer了30个数字，注意，是30，29，28...的输出，不是30个30。这一点验证了，defer语句相关变量的计算会在定义的地方执行。
*/
func tryDefer() {
	defer fmt.Println("defer1")
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print too many")
		}
	}
}

/**
go语言中利用defer实现资源管理的例子，体会go语言的设计理念
项文件中写入斐波那契数列的n项:
1. 创建/打开文件和关闭文件是成对出现的。所以在创建问价成功之后，就直接defer调用Close()而不用等到中间各种处理之后，或者不同情况之后，再去调用Close()
就是这样的流程：我创建了一个文件，然后告诉go语言：go语言，我做完事之后你记得帮我关闭（defer调用Close()）
2. 利用bufio的缓存写入功能：创建一片缓存，然后向缓存写入数，什么时候缓存满了，就直接把缓存的内容写入文件，然后继续这个过程。最后，需要考虑需要释放缓存。
使用defer之后，就是这样的流程：我创建了一片缓存，然后告诉go语言：go语言，我做完事之后，你记得帮我释放缓存（defer调用Flush）


所以这段代码从直接按照人类的逻辑思维去书写就可以：
1. 打开一个文件，打开成功之后，心想，一会要记得关闭这个文件（defer file.Close()）。
2. 利用刚刚打开的文件，通过缓存写入的功能，去写如数据。
3. 创建一片缓存，然后向缓存写入数，什么时候缓存满了，就直接把缓存的内容写入文件，然后继续这个过程。最后，需要考虑需要释放缓存（defer w.Flush()）。

整个过程，不用想其他语言那样去考虑，什么时候释放缓存，什么时候关闭文件。先释放缓存还是先关闭文件。你要做的就是需要按照正常的逻辑去处理：打开了资源，我就要关闭资源。
关闭资源直接用defer去写，执行完成之后，go语言会帮我执行defer。
*/
func writeFile(filename string, n int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	defer w.Flush()
	f := fib.Fibonacci()
	for i := 0; i < n; i++ {
		fmt.Fprintln(w, f())
	}
}

/**
错误处理
if err != nil 的处理。以打开文件为例。步骤
1. 判断err，是哪种错误。需要去返回err的方法中去找，看看注释，看看return，然后确定。
2. 得到err是a类型的错误，做一个类型断言之后。单独处理。打log。输出，等等。
3. 如果返回的err不是预测的错误。则另想办法。
errors这个包里面有生成自定义错误的方式。
*/

func myError() {
	_, err := os.Open("errorHandling/abc.txt")
	if err != nil {
		//s := errors.New("操作异常")
		//panic(s)
		switch e := err.(type) {
		case *os.PathError:
			fmt.Printf("Error: when %s the file %s , error( %s ) happend.", e.Op, e.Path, e.Err)
		default:
			e = errors.New("操作异常")
			panic("e")
		}
	}
}

func main() {
	//tryDefer()
	myError()
	writeFile("errorHandling/fibonacci.md", 40)
}
