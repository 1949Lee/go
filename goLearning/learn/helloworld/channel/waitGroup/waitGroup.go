package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	//var c chan int // 这种方式虽然定义了channel c，但是 c == nil。并不能使用
	for n := range w.in {
		fmt.Printf("Worker%d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, w *sync.WaitGroup) worker {
	_w := worker{
		in:   make(chan int),
		done: func() {
			w.Done()
		},
	}
	go doWorker(id, _w)
	return _w
}


// 利用系统自带的包来实现：并发执行两批任务，发送一批小写，发送一批大写。然后等待接受完成。
func channelWaitGroup() {
	var wait sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wait)
	}

	wait.Add(len(workers) * 2)
	for i := 0; i < len(workers); i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < len(workers); i++ {
		workers[i].in <- 'A' + i
	}

	wait.Wait()

}

func main() {
	channelWaitGroup()
}
