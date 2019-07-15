package main

import "fmt"

func printSLice(s []int) {
	fmt.Printf("slice = %v, len = %d, cap = %d\n******************\n", s, len(s), cap(s))
}

/*添加slice的处理*/
func appendSlice() {
	arr := [...]int{0, 2, 4, 6, 8, 10}
	s1 := arr[:4]
	//***********slice的添加删除等操作
	s2 := append(s1, 12)
	fmt.Println("s2 = ", s2)
	s3 := append(s2, []int{14, 15}...) // 某次添加超出切片的原数组的时候，append后将会产生新的数组，切片是这个新数组的切片。请旧数组注意最后一个为10
	fmt.Println("arr = ", arr)
	fmt.Println("s3 = ", s3)
	s4 := append(s2, 16) // 请注意最后一个为16
	fmt.Println("s5 = ", s4)
	fmt.Println("arr = ", arr)
}

/*创建slice，及slice动态添加时的处理*/
func createSlice() {

	// 创建方式一：简单创建
	var s []int // 此时已经有初始值：nil
	for i := 0; i < 10; i++ {
		printSLice(s)
		s = append(s, 2*i+1) // 当添加时容量不够时，容量变为原来的2倍。
	}
	printSLice(s)

	// 创建方式二：赋值创建
	s2 := []int{2, 4, 6, 8, 10}
	printSLice(s2)

	// 创建方式三：明确长度创建
	s3 := make([]int, 16)
	printSLice(s3)

	// 创建方式四：明确长度、明确容量创建
	s4 := make([]int, 16, 32)
	printSLice(s4)
}

//复制slice
func copySlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := make([]int, 16)
	copy(s2, s1)
	printSLice(s2)
}

//删除slice中某一项
func deleteSlice(i int) {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := append(s1[:i], s1[i+1:]...)
	printSLice(s2)
}

func popSLice(isFront bool) (int, []int) {
	s1 := []int{2, 4, 6, 8, 10}
	if isFront {
		front := s1[0]
		s2 := s1[1:]
		printSLice(s2)
		return front, s2
	} else {
		tail := s1[len(s1)-1]
		s2 := s1[:len(s1)-1]
		printSLice(s2)
		return tail, s2
	}
}

func main() {
	//appendSlice()
	//createSlice()
	//copySlice()
	//deleteSlice(1)
	//popSLice(true)
	//popSLice(false)
	arr := []int{2, 4, 6, 8}
	s1 := arr[:3]
	s1[1] = -1
	s2 := s1[1:]
	fmt.Println(arr, s1, s2)
}
