package main

import "fmt"

func updateSlice(s []int) []int {
	s[0] = -1
	return s
}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	//***********改变切片
	fmt.Println("updateSlice(s1)")
	updateSlice(arr[:])
	fmt.Println("arr = ", arr)
	fmt.Println("arr[:] = ", arr[:])

	//***********对slice的slice
	s2 := arr[:6]
	s2 = s2[2:]
	fmt.Println("slice的slice：", s2)
	arr2 := [...]int{0, 2, 4, 6, 8, 10}
	s1 := arr2[:4] // 虽然你只能看到s1的0至3，但是其内部还有隐藏的4至5。就是说s1的长度是4，容量是6
	// 虽然切片还有隐藏的部分，但是不能通过下标直接访问s1的隐藏部分，s1[4]、s1[5]会报错
	ss := s1[3:6] // 所以ss能够再切出s1隐藏的部分
	fmt.Printf("s1: %v *** length:%d *** capacity:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("ss: %v *** length:%d *** capacity:%d\n", ss, len(ss), cap(ss))
}
