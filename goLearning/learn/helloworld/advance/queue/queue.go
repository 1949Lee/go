/*
这是一个先进先出队列的包
*/
package queue

// 表示队列，队列本身是一个int切片
type Queue []int

// 队列的入队方法，将传入的int压入队列
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 队列的出队方法，将队列的首部第一个元素弹出，方法将返回弹出的值
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// 判断队列是否为空的方法。若队列为空，则返回true；否则返回false
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
