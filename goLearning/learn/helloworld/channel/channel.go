package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//var c chan int // 这种方式虽然定义了channel c，但是 c == nil。并不能使用
	//判断这个channel是否关闭，方法一
	for n := range c {
		fmt.Printf("Worker%d received %c\n", id, n)
	}
	////判断这个channel是否关闭，方法二
	//for {
	//	n, ok := <-c
	//	if!ok {
	//		break
	//	}
	//	fmt.Printf("Worker%d received %c\n", id, n)
	//}
}

func createrWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelDemo(){
	var chans [10]chan int
	for i := 0; i < 10; i++ {
		chans[i] = createrWorker(i)
		//go worker(i, chans[i])
	}

	for i := 0; i < len(chans); i++ {
		chans[i] <- 'a' + i
	}
}

func bufferedChannel() {
	//设置channel有一个大小为3的缓冲区。可以存储3个数据。当缓冲区没有空间时，需要goroutine来接受数据
	c := make(chan int, 3)
	go worker(11, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}
func channelClose() {
	c := make(chan int)
	go worker(11, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	close(c) // 关闭channel之后，接收方要要判断这个channel是否关闭，有两个方法。
}

func main() {
	channelDemo()

	//bufferedChannel()

	//channelClose()
	time.Sleep(time.Millisecond)

}
