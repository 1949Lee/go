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


### tree/traverse.go中有获取输的最大深度的非递归方式（深度和广度两种）

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
[文章1](https://segmentfault.com/a/1190000019389694?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com#articleHeader8)
[文章2](https://www.jianshu.com/p/84d231048bc4)


### 测试
go help testflag可以查看能运行的命令
[文章1](http://c.biancheng.net/view/124.html)
1. 普通测试：详见maps_test.go; 可用命令行执行
2. 覆盖率测试：详见maps_test.go; go内置了覆盖率测试的语句。可用命令行执行
3. 性能测试：详见maps_test.go; go内置了性能能测试的语句。可用命令行执行

PProf:GO 性能测试工具PProf
[文章1](https://blog.csdn.net/guyan0319/article/details/85007181)
[文章2](https://www.jianshu.com/p/4e4ff6be6af9)

代码性能优化过程：
1. 输出性能测试时cpu使用数据。go test -bench . -cpuprofile=cpu.out
2. 查看cpu使用数据。go tool pprof。// 用web查看
3. 根据生成的svg查看慢在哪里，然后优化代码。
4. 继续以上3个步骤，直至代码优化到满意为止（总运行时间越来越快，直至满意）
5. 注意空间换时间的问题，是否值得。

webApi测试可以通过httptest包来测试：利用包来启动一个服务，然后发送请求。详见errorhandling_test.go


### 文档
godoc表示生成go文件的注释，go doc表示查看go文件的注释。
godoc的举例使用请看queue.go


### go语言的并发模型：CSP

### 协程
进程、线程、协程。协程并不是go特有的，只不过go支持的比较好。
协程特点：
1. 非抢占式多任务处理，操作系统不会像线程那样，去循环的切换线程的执行。有协程主动交出控制权
2. 编译器、解释器。虚拟机层面的多任务。go语言有背后有自己的调度器。
3. 多个协程可能在一个或多个线程上运行。这点由调度器决定。
4. 子程序是协程的一个特例
实现：任何函数加上go执行，就会将函数送给调度器运行。调度器会在合适的点进行切换（不是随意切换）

goroutine可能会切换的点
+ I/O,select
+ channel
+ 等待锁
+ 函数调用（有时）
+ runtime.Gosched()
以上只是参考。


### channel
定义一个channel`c := make(chan int)`
channel发送数据之后一定要有人来接受
channel作为返回值得时候有三种类型。
1. func() chan int {} 这种函数返回的channel可以发送也可以接受。
2. func() chan<- int {} 这种函数返回的channel只可以向这个channel发送数据。
3. func() <-chan int {} 这种函数返回的channel只可以从这个channel接受数据。
例子：利用channel实现树的遍历。详见advance/tree/traverse.go


### 传统的同步机制(尽量不去使用)
除了go比较特色的channel+go routine之外，go也可以像传统语言一样的方式去并发处理
1. Mutex、atomic：go内置的mutex互斥变量和atomic操作。
2. Cond:

### http、http.Client
http.Client可以构造一个客户端，然后通过client.Do(req)的方式来发送请求。
可以使用[net/http/pprof](https://cloud.tencent.com/developer/section/1143647)来查看web服务的性能
查看相关资料：使用http.FileServer来提供静态内容，css、js、图片、index.html等。


### html/template 模板引擎
查看相关资料：语法+测试



### 爬虫实战部分中可以记录的内容。
1. 获取URL网页之后根据不同的编码格式（gbk等）统一为utf8编码格式：详见crwaler/main.go中`determineEncoding`方法
2. golang中的[jQuery](https://github.com/PuerkitoBio/goquery)。
3. URL实现去重，即当要爬取新的URL的时候，如何确定这个URL已经被爬取过了。需要将爬取过的URL存储下来，然后新的URL去存储里比对，看看有没有爬取过。所以存储的方式有：
    + 哈希表(go中的map)
    + 计算MD5等值哈希，再存哈希表
    + 使用bloom filter多重哈希结构
    + 使用redis等key-value存储系统实现分布式去重

### pprof+wrk测试爬虫实战部分性能

1. 自动发送请求，使用工具wrk
	wrk -t1 -c1 -d30s --script=wrkPost.lua http://127.0.0.1:1314
2. 监听cpu使用情况，使用go自带的pprof。命令执行后会监听一段时间（默认30s）内
的端口使用情况。可以手点，也可以用步骤1的工具	go tool pprof http://127.0.0.1:1314/debug/pprof/profile
输入help可查看具体命令如下：
Commands:
    + callgrind        Outputs a graph in callgrind format
    + comments         Output all profile comments
    + disasm           Output assembly listings annotated with samples
    + dot              Outputs a graph in DOT format
    + eog              Visualize graph through eog
    + evince           Visualize graph through evince
    + gif              Outputs a graph image in GIF format
    + gv               Visualize graph through gv
    + kcachegrind      Visualize report in KCachegrind
    + list             Output annotated source for functions matching regexp
    + pdf              Outputs a graph in PDF format
    + peek             Output callers/callees of functions matching regexp
    + png              Outputs a graph image in PNG format
    + proto            Outputs the profile in compressed protobuf format
    + ps               Outputs a graph in PS format
    + raw              Outputs a text representation of the raw profile
    + svg              Outputs a graph in SVG format
    + tags             Outputs all tags in the profile
    + text             Outputs top entries in text form
    + top              Outputs top entries in text form
    + topproto         Outputs top entries in compressed protobuf format
    + traces           Outputs all profile samples in text form
    + tree             Outputs a text rendering of call graph
    + web              Visualize graph through web browser
    + weblist          Display annotated source in a web browser
    + o/options        List options and their current values
    + quit/exit/^D     Exit pprof




3. 输入proto后回车。会在当前目录下生成监听的一个结果（一压缩文件）
4. 利用刚刚生成的压缩文件，可以在线查看：
	go tool pprof -http=:8080 profile001.pb.gz
	
	
### ElasticSearch 初识
+ <server>:9200/index/type/id
 <br>index相当于关系型数据库中的database
 type相当于关系型数据库中的表
+ 不需要预先创建
+ type中数据类型可以不一致
+ 可以使用_mapping来配置类型
+ 使用REST接口
+ 使用put或post创建或修改数据，使用post可以省略id
+ 用get获取各种数据。
+ Get <index>/<type>/_search?q=全文搜索关键。来全文搜索
