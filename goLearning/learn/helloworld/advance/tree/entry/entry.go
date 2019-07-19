package main

import (
	"fmt"
	"goLearning/learn/helloworld/advance/tree"
)

//根据tree包里的node扩展
type MyTreeNode struct {
	node *tree.Node
}

//根据tree包里的node扩展
func (myNode *MyTreeNode) posOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.posOrder()
	right := MyTreeNode{myNode.node.Right}
	right.posOrder()
	myNode.node.Print()
}

func main() {

	//定义方式一
	var root tree.Node

	//定义方式二
	//root := treeNode{}

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{Value: 1}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Left.SetValue(4)
	root.Left.Right = tree.CreateNode(10)

	root.Print()
	fmt.Println()

	//结构体调用自己的方法时不用关心是指针调用该方法，还是结构体调用该方法，更不用关心该方法定义的是指针的还是结构体的，go的编译器会自动识别
	root.Left.Right.SetValue(2)
	root.Left.Right.Print()
	fmt.Println()
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count of root:", nodeCount)

	posRoot := MyTreeNode{node: &root}
	posRoot.posOrder()
	fmt.Println()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{Right: nil, Value: 6},
	}

	fmt.Println(root, nodes)
}
