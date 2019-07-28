package main

import (
	"fmt"
	"time"
)

func main() {
	//a := [10]int{}
	//for i := 0; i < 1000; i++ {
	//
	//	// goroutine的传入参数相当于将协程的上下文保存下来。如果不传，goroutine可能会报错（变量跳出循环时时1000，a[1000]报错了）。
	//	go func(i int) {
	//		for {
	//			a[i]++
	//			runtime.Gosched() // 协程主动让出执行控制权。不进行手动让出的话，第一个协程会一直执行a[i]++。然后就死机了。
	//		}
	//	}(i)
	//}

	for i := 0; i < 100000 ; i++ {
		go func(i int) {  // goroutine要使用for死循环。
			for {
				fmt.Printf("hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)


	/*
	去掉这句话会什么也没输出。因为go func是并发的执行。main程序很快循环完1000次，开了1000个协程并发执行。但是每个协程还没来得及执行具体的语句，main函数就执行完，go程序就退出了
	*/
	//time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
