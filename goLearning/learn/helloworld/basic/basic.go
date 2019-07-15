package main

import (
	"fmt"
	"math"
)

// 这里的变量的作用于是包内变量，并不是全局变量，go语言不存在全局变量
var aa = 3
var _ss = "kkk"

// 变量集中定义
var (
	cc = 5
	dd = true
)

/* 错误写法
bb := true
*/

func variableZeroValue() {
	// 变量定义时，go会附一个初始值
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	// 多个相同类型的变量定义
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// ***** 推荐写法
func variableTypeDeduction() {
	// 变量定义自动判断类型
	var a, b, c, s = 3, 4, true, "def"
	var (
		abc = 1
		bcd = "bcd"
	)
	fmt.Println(a, b, c, s, abc, bcd)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 勾股定理
func triangle() {
	a, b := 3, 4
	var c int
	// math.Sqrt方法需要float64的参数，方法的返回也是float64，所以需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("%d \n", c)
}

// 常量
func consts() {
	const filename = "note.md"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Printf("%d \n", c)
}

// 枚举
func enum() {

	//普通枚举
	const (
		c      = 1
		python = 2
		java   = 3
	)
	fmt.Println(c, python, java)

	//自增值枚举,iota默认从0开始
	const (
		_aa = iota
		_   // 表示跳过这次自增值
		bb  // bb = 2
		cc
		dd
	)
	fmt.Println(_aa, bb, cc, dd)

	// iota的高级用法：文件大小，b,kb,mb,gb,tb,pb
	const (
		b  = 1 << (10 * iota) // 1字节
		kb                    // 1024字节
		mb                    // 1024kb
		gb                    // 1024mb
		tb
		pb
	)
	fmt.Println(b, kb, mb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, _ss, cc, dd)
	triangle()
	a := 4.99
	fmt.Println(int(a))
	consts()
	enum()
}
