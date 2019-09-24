package tree

import "fmt"

//遍历节点
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

// 传递一个方法遍历树
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

// 通过channel的方式，来遍历树（可以通过channel输出的遍历节点做自己的事），效果上来讲和TraverseFunc（回调函数的方式）遍历有异曲同工之妙
func (node *Node) TraverseByChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}

// 非递归的方式求树的深度，广度优先遍历tree
func (node *Node) MaxDepth() int {
	depth := 0
	queue := make([]*Node, 0)
	if node != nil {
		queue = append(queue, node)
	}

	for len(queue) > 0 {
		depth++
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return depth
}
