package main

import "fmt"

//函数中的sum就是闭包。当b函数作为a函数的返回值时，a函数返回的不仅仅是一个函数b，还有b中所需要运行的所有变量依赖
func adder(i int) func(int) int {
	sum := i
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder(0)
	for i := 0; i <= 10; i++ {
		fmt.Printf("0 + 1 + ... +%d = %d \n", i, a(i))
	}
}
