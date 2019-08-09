package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)


// 利用sync.Mutex实现变量的原子操作。这里仅做演示，实际应用中直接使用go内置的atomic操作。
type atomicType struct {
	value int
	lock sync.Mutex
}

func (a *atomicType) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++


	// 利用匿名函数实现代码块的原子操作
	/*
	func () {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
		// other operations...
	}()
	*/
}


func (a *atomicType) get() int{
	a.lock.Lock()
	a.lock.Unlock()
	return a.value
}

func main() {

	a := atomicType{}
	a.increment()
	go func() {
		a.increment()
	}()
	var b int32 = 1

	// go内置的int增加的原子操作。
	atomic.AddInt32(&b, 2)
	time.Sleep(time.Millisecond)

	fmt.Println(a.get())
	fmt.Println(b)

}
