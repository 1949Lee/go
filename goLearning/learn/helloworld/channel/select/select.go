package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker%d received %d\n", id, n)
	}
}

func createrWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var values []int
	var w = createrWorker(0)
	//var ticker = time.Tick(time.Second)
	var life = time.After(10 * time.Second)
	for {
		var activeWorker chan int
		var activeValue = 0
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("received time out: ")
		//case <-ticker:
		//	fmt.Println("length of value queue: ", len(values))
		case activeWorker <- activeValue:
			values = values[1:]
		case <-life:
			fmt.Println("The End")
			return
		}
	}
}
