package queue

import "fmt"

func ExampleQueue_Pop() {
	q := Queue{1,2,3}
	q.Push(4) // 将4压入队列
	fmt.Println("origin queue is :",q)
	fmt.Printf("head first of the queue:%d, has been popped\n",q.Pop())
	fmt.Println("the queue after pop action is :",q)

	// Output:
	// origin queue is : [1 2 3 4]
	// head first of the queue:1, has been popped
	// the queue after pop action is : [2 3 4]
}
