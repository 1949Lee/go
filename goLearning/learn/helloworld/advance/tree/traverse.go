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

/* 非递归的方式求树的深度，广度优先遍历tree。
LeetCode的题，求树的最大深度：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/submissions/
*/
func (node *Node) MaxDepth() int {
	depth := 0
	queue := make([]*Node, 0)
	if node != nil {
		queue = append(queue, node)
	}

	// 每次取出队列中的全部，将取出节点的子节点加入队列（即下一层所有节点）
	for len(queue) > 0 {
		depth++
		length := len(queue)
		// 一次性取出队列中的全部，这些都是同一层的兄弟节点。
		for i := 0; i < length; i++ {

			// 这一层的每个节点都要将该节点的所有子节点（即下一层所有节点）加入队列。
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// 弹出本层所有节点的操作（留下剩余节点）
		queue = queue[length:]
	}
	return depth
}

//func maxDepth(root *TreeNode) int {
//	depth := 0
//	stack := make([]*TreeNode,0)
//	if root != nil {
//		stack = append(stack, root)
//	}
//
//	for len(stack) > 0 {
//		top := stack[len(stack) - 1]
//		if top.Left != nil {
//			stack = append(stack,top.Left)
//			continue
//		}
//		if top.Right != nil {
//			stack = append(stack,top.Right)
//			continue
//		}
//		if depth < len(stack) {
//			depth = len(stack)
//		}
//		stack = stack[:len(stack) - 1]
//	}
//	return depth
//}
