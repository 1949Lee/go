package main

import (
	"fmt"
)

func doWorker(id int, c chan int, done chan bool) {
	//var c chan int // 这种方式虽然定义了channel c，但是 c == nil。并不能使用
	for n := range c {
		fmt.Printf("Worker%d received %c\n", id, n)
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}


// 顺序发送一批小写，然后接受。然后在顺序发送一批大写。然后执行。
func channelDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < len(workers); i++ {
		workers[i].in <- 'a' + i
	}

	for _,worker := range  workers {
		<-worker.done
	}

	for i := 0; i < len(workers); i++ {
		workers[i].in <- 'A' + i
	}

	for _,worker := range  workers {
		<-worker.done
	}
}

func main() {
	channelDemo()
}
