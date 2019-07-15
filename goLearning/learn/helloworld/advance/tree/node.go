package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

// treeNode的方法-结构体接受者，函数名前面定义了一个接受者，表明该函数是结构体treeNode的一个方法，固俗称方法，下面定义的方法是值传递，无法改变接受者的内容
func (node Node) Print() {
	fmt.Print(node.Value)
}

// treeNode的方法-结构体指针接受者，定义指针接受者才可以改变接受者的数据（引用传递），这种引用传递时主意判断指针是不是nil
func (node *Node) SetValue(value int) {
	if node == nil {
		return
	}
	node.Value = value
}

//工厂函数，其他语言中的构造函数。注意，返回的是局部变量的内存地址
func CreateNode(value int) *Node {
	//函数中返回结构体的指针，这个结构体会被创建在堆上。不同于c++。因为go中这样做，是的，这个被返回的结构体可以在函数外部使用
	return &Node{Value: value}
}
