package main

import (
	"fmt"
	"goLearning/learn/helloworld/advance/retriever/retriever"
	"time"
)

// 接口，该接口内部有一个get方法，只要实现了这个get方法的结构体都视为实现了这个接口
type Retriever interface {
	Get(url string) string
}

// download方法需要一个实现者，这个实现者就是上面那个接口。不管这个接口背后是什么结构体
func download(r Retriever, url string) string {
	return r.Get(url)
}

func main() {
	var r Retriever
	r = &retriever.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(download(r, "https://www.jiaxuanlee.com"))

	// 接口是由结构体实现的，那接口到底是什么样的:这个接口的类型是retriever.Retriever，这个接口的值是{Mozilla/5.0 1m0s}
	fmt.Printf("%T， %v\n", r, r)

	/*结论：假设你要组建一直军队，要征兵，然后你指定了一套标准：
	1.年龄18-40岁。
	2.能徒手劈碎2块砖。
	就这两个要求。那么所有复符合这个标准的人其实都是你要找的兵。我不管你来自哪，怎么训练自己的。只要复合标准，都可以当我的兵。
	当然这套标准的要求其实可以不止两个。
	那这个例子对应到go中的接口，就是：使用者规定了一套标准，这套标准包含一个或多个方法。只要实现了这套标准的结构体，都是我需要的结构体（实现者）。那么这套标准就是一个接口。
	实现了某个使用者规定的接口的实现者就可以作为这个使用者的实现者*/

	//接口的内部会规定一些方法，实现者去实现这个接口就需要实现这些方法。即定义一个结构体，然后该结构体定义一些方法，该方法正好是一些接口规定的，那么这个结构体就实现了对应的接口。
	//所以实现接口的本质就是结构体实现方法。所以接口的方法就有值传递（结构体本身）和指针（结构体的指针）两种方式。

	//可以用类型断言来确定该接口的类型
	//方法一 接口r是否是某种结构体或者结构提的指针，若是，则返回该结构体
	switch v := r.(type) {
	case *retriever.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	//方法二 接口r是否符合某个接口，若是，则返回该接口的r实现
	fmt.Println("接口的方法地址:", r.(Retriever).Get)

	//类型断言的成功和失败
	if v, ok := r.(*retriever.Retriever); ok {
		fmt.Println("TimeOut:", v.TimeOut)
	} else {
		//转换出错了 fmt.Errorf("error %s:", v)
	}

}
