### 变量定义
见 basic.go
变量常量的标识符慎用大小写，因为表示不同的含义。大写开头表示public类型


### 内置变量类型
* bool, string 
* (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
  * uint表示无符号整数，int表示有符号整数
  * int和uint并不是固定的长度，他根据CPU来决定，32位的CPU，则int是32位；64位CPU，则int是64位
  * uintptr表示指针
* byte, rune
  * byte是8位一字节
  * rune是字符类型（其他语言中的char），用32位（4个字节来表示），因为字符编码集（utf-8），有的字符（表情）是3个字节的，所以go直接用4个字节来存储字符。
* float32, float64, complex64, complex128
  * complex64和complex128表示数学中的复数（z=a+bi），复数分为实部和虚部，各站一般。即：complex64的复数，实部占32位，虚部占32位。
  
### 强制类型转换
```
	a, b := 3, 4
	var c int
	// math.Sqrt方法需要float64的参数，方法的返回也是float64，所以需要强制转换
	c = int(math.Sqrt(float64(a * a + b * b)))
	// int(1.9)则将1.1强制转化为int类型 1， 浮点转int会直接取整数部分，而非四舍五入
```

### 常量与枚举
常量
普通枚举，自增值枚举


### 控制语句
1. if语句：见control.go
2. switch语句：见control.go
3. 循环语句：见loop.go
* go语言没有while，for的三个语句，分别都能省略！！！，只有中间终止条件的时候，就是while。三个条件都省略的时候，就是死循环，详见loop.go

### 函数
见func.go
* 函数的参数的传递的方式是值传递，只有值传递，想要实现引用传递，一些内置类型，就需要将参数定义成对应的指针类型，然后调用时传递参数地址。
* 自定义类型需要做好封装，然后确定需要值传递还是引用传递

### 指针
见pointer.go

### 数组
GO中数组是值类型！GO中数组是值类型！GO中数组是值类型！
var a [5]int 是数组
var a []int 是切片
数组作为参数传递就会拷贝数组，需要特别注意，所以可以利用数组指针来实现传递参数，所以很麻烦啊。
所以额GO中不直接使用数组，而是使用切片
详见arrays.go


### 切片