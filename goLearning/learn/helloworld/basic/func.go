package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		// 当函数返回多个值得时候，可以用'_'来忽略函数的某些返回值，参考自增值常量中的'_'
		q, _, _ := div(a, b)
		return q
	default:
		panic("unsupported operation:" + op)
	}
}

// 函数多个返回值得推荐写法
func div(a, b int) (int, int, error) {
	if b < 0 {
		return 0, 0, fmt.Errorf("The divisor must not be less than 0")
	} else {
		return a / b, a % b, nil
	}
}

//简写形式，可以给返回值起名，但是需要慎用，因为函数体一旦变长你很可能不知道最后q和r在哪里赋过值
func divNamed(a, b int) (q, r int) {
	q, r = a/b, a%b
	return
}

//函数作为参数
func apply(op func(int, int) int, a, b int) int {

	//反射获取函数指针
	p := reflect.ValueOf(op).Pointer()

	//跟胡函数指针获取函数名并打印
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//函数参数的不定列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func closure() []func() {
	var q []func()
	for i := 0; i < 10; i++ {
		//正确使用上下文及闭包
		//func (i int) {
		//	q = append(q, func() {
		//		fmt.Println(i)
		//	})
		//}(i)

		//// 错误使用上下文及闭包
		//q = append(q, func() {
		//	fmt.Println(i)
		//})
	}
	return q
}

func main() {
	//fmt.Println(eval(1, 3, "*"))
	//fmt.Println(eval(6, 3, "/"))
	//if q, r, err := div(13, 4); err == nil {
	//	fmt.Println(q, r)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//// 使用匿名函数
	//fmt.Println(apply(
	//	func(a int, b int) int {
	//		return int(math.Pow(float64(a), float64(b)))
	//	}, 3, 4))
	//fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// 争取和错误使用上下文及闭包
	q := closure()
	q[0]() // 并没有打印出i = 0
	q[1]() // 并没有打印出i = 1
}
