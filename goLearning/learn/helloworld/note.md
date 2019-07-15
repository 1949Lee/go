### 变量定义
见 main.go
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
```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7} // 原数组
arr[2:],arr[:] //都是切片
```
改变切片里的值就会改变原数组，改变原数组的值相应的切片就会改变。
切片作为函数参数的时候，改变切片的内容就会改变原数组。
所以切片是对数组的真正数据存储的映射
切片对切片的扩展：详见slices.go、slicesOperations.go

### Map
键值对集合，
```go
map[string]string{} // 表示一个key是string类型，value是string类型的。map。
map[string]map[int]int{} // 复合map：表示一个key是string类型，value是另外一种map。
```
遍历map,每次遍历输出的顺序可能会不同，因为键值对在map中是无序的，go中的map是哈希map
map中的key：1、必须是可以比较的类型；2、除了slice、map、function的内置类型；3、不含slice、map、function的Struct


### 字符串及rune的处理
go中的char类型就是rune，详见string.go


### 结构体和方法
结构体的方法的接受者可以定义两种：结构体本身的类型，和结构体的指针。无论定义哪一种，调用结构体的方法的时候即可以使用结构体的指针，也可以使用结构体。go的编译器会自动根据结构的方法定义的接受者类型来做对应。
详见node.go

### 包和封装
封装：
1. 名字一般使用驼峰CamelCase
2. 包内的方法或者变量——首字母大写：public；首字母小写：private

包：
1. 每个目录一个包，可以和目录名字不一样
2. main包包含可执行入口。如果目录下有main函数，那该目录下只能是main包。
3. 为结构定义的方法必须放在同一个包内，但是可以放不同的文件。

扩展别人的包或者扩展系统的包

例子见advance/tree

### 内存分配是go自助管理的，需要理解
[1](https://segmentfault.com/a/1190000019389694?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com#articleHeader8)
[2](https://www.jianshu.com/p/84d231048bc4)