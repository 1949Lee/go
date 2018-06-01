---
title: Mac+VSCode+Go开发环境搭建
tags: 开发环境搭建,Go
---
## 解决的问题
使用`vscode`编辑器编写`go`，然后进行编译、调试（打断点）。
## 安装Go
搜索官网，下载安装包自行安装。此处不详细说了。安装时请记住你的安装路径。默认是`/usr/local/go`。
安装完成后，需要配置一个环境变量。
如果你的终端用的zsh（iterm2）。那就需要到`~/.zshrc`文件里添加。
如果你用的MAC默认的终端，那就需要到`~/.bash_profile`文件里添加。
```
export GOPATH=/usr/local/go/bin
export PATH=$PATH:$GOPATH
```
添加完以上代码后需要执行`source ~/.zshrc`或者`source ~/.bash_profile`命令，使配置生效。然后重新打开一个命令行终端。
输入以下代码`go version`。输出版本则证明基本成功。
然后我们编写一个hello.go文件测试以下。
1. 新建一个hello.go文件，然后打开并输入以下代码
``` golang?linenums
package main

import "fmt"

func main() {
	fmt.Print("Hello World\n")
}
```
 2. 在命令行终端中，进入到hello.go所在的路径，然后输入以下命令并回车执行。
```shell
go run hello.go
```
会看到输出Hello World。
到此go就算安装完成了。但是我们不肯能这么编写go代码，这效率太低了。
## 集成开发环境IDE:Microsoft Visual Studio Code
个人比较喜欢VSCode，搜索官网，下载安装包自行安装。继续跳过。
## 完美的Go开发环境
为什么VSCode会火，因为插件丰富，然后微软做这种编辑器就是有经验啊。
我们需要下载一个Go的插件，这个插件会让我们有智能提示，语法错误检查（lint）等功能。十分强大，也是微软官方出的。
建议大家先去插件的介绍页看看
[Go For Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
这个插件用到了一些go官方和开发者的依赖，所以你需要安装这些依赖。
这些依赖需要安装到你的go路径下（安装时的路径：默认`/usr/local/go/bin`）。我安装时需要管理员权限。
我们用到了这些依赖（其实插件的介绍页面里已经写了网址）
[需要安装的依赖](https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on)
在链接的最后已经列出来了
``` shell
# 2018-06-01日的版本用到了这些
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/tools/cmd/godoc
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/goimports
go get -u -v github.com/cweill/gotests/...
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/josharian/impl
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
```
VSCode提供了安装所有的快捷通道，但是需要sudo（管理员权限），所以我们分条执行
 1. 首先进入到Go语言的bin目录下（安装位置），默认`/usr/local/go/bin`。
 2. 然后咱么开始安装，就是在我上面列出的`go`前面，加上sudo（其实就第一次用），然后输入开机密码，你会看到安装的过程。然后就看到在`/usr/local/go/bin`路径下出现了我们需要的依赖。
 
 
 其实到此为止，我们就可以进项开发了。。。。。
 但是但是。我们调试怎么办，我想打断点。我想看变量啊。我想看调用堆栈。
 ## 调试的安装
 我们的VSCode插件（上面提到的），其实也提到了。
 你仔细观察这个插件安装的依赖，会发现很有特色，github.com的依赖放到了一起。调试依赖官方的安装说明并没有生效。所以我才发现了这个特点。所以我们只要把调试依赖的github代码像其它依赖一样安装好，就可以在VSCode中使用了呀。
 所以流程是这样的（一下已默认安装路径为例，否则需要替换所有的`/usr/local/go/bin`）:
 1. 建立相应的目录。到`/usr/local/go/bin/github.com/`路径下建立这样的路径,建立之后应该存在这样的路径`/usr/local/go/bin/github.com/derekparker/delve`。
 2. 下载调试依赖的git库：`https://github.com/derekparker/delve`需直接zip包，然后把源代码解压，然后把源代码直接放到上面的路径下（上面的路径下应该直接是源代码，可以找到Makefile文件！！！！！！）。
 3. 执行命令`make install`。
 4. 配置VSCode的launch.json文，具体怎么配置给个官网链接，英文的自己去看，看不懂你根本不配看这个文章。`https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code`。
 5. 打断点进行测试（测试的项目可以是我们一开始编写的hello.go文件）。