package main

import "fmt"

//数组作为参数传递就会拷贝数组，需要特别注意，所以可以利用数组指针来实现传递参数
func printArray(a [5]int) {
	a[0] = 100
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{3, 4, 5}

	// 自动定长数组
	arr3 := [...]int{1, 4, 9, 16, 25}
	var grid [2][3]bool

	printArray(arr1)
	printArray(arr3)

	for i, v := range arr3 {
		//for _,v := range arr3 { // 只获取数组元素，不获取下标
		fmt.Println(i, v)
	}

	fmt.Println(arr1, arr2, arr3, grid)
}
