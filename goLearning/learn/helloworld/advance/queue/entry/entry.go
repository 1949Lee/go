package main

import (
	"fmt"
	"goLearning/learn/helloworld/advance/queue"
)

func main() {
	q := queue.Queue{-1, 2, 3, 4}
	q.Push(5)
	q.Push(6)
	front := q.Pop()
	fmt.Println(front)

	for !q.IsEmpty() {
		i := q.Pop()
		fmt.Println(i)
	}
}
