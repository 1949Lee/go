package parser

import "testing"

//性能测试，结果注意单位，ns表示纳秒
func BenchmarkParse(t *testing.B) {
	md := `~s*打双打a~~打算sda1*das哒2*~~
这是一个连接：[李佳轩的个人博客——*镜中之人*](http://www.jiaxuanlee.com)
# [JDB-H5-UI-VUE](http://git.jdb-dev.com/zen/jdb-h5-ui)

### 首次使用
仓库克隆到本地以后，首先执行。这条命令会做一些必要的准备工作。
然后你再进行常规的等等操作。所有npm命令的具体含义可以查看本文档的npm命令部分

### 开发组件库的组件
确保自己已经在项目根目录
1. 在src目录下新建自己的组件目录（如button）。
2. 在刚刚新建的组件目录下新建文件index.js。js的内容形如（button为例）：

3. 在组件目录下新建一个src目录，在src目录下新建对应组件的vue文件。然后开始开发组件。

开发组件的时候可能需要调试或者开发完成后想写一个demo试试自己开发的组件。可以参考本文档的调试组件部分。

### 调试组件
sites目录下是一个整的vue项目。如果想要调试组件。可以修改sites/debug/index.vue文件的内容
比如要调试button组件。sites/debug/index.vue文件中可以引入src目录下的button组件来作为自己的一个组件去调试，即：
写好之后，可以运行查看效果。所有npm命令的具体含义可以查看本文档的npm命令部分

### 打包组件库`
	//ans := 9
	t.ResetTimer()

	//性能测试循环测试的次数，由go test自动计算（t.N）
	for i := 0; i < t.N; i++ {
		l := Line{Origin: []rune(md), Tokens: []Token{}}
		l.Parse()
	}
	//if actual := l.p; actual != ans {
	//    t.Errorf("input string %s got %d; expected %d.", s, actual, ans)
	//}
}
