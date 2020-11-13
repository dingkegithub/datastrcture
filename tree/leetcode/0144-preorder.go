package leetcode

import "container/list"

//
// 解法1： 根-左-右
//
func preorderTraversal(root *TreeNode) []int {
	out := make([]int, 0, 10)
	preOrder(root, &out)
	return out
}

func preOrder(node *TreeNode, out *[]int) {
	if node == nil {
		return
	}
	*out = append(*out, node.Val)
	preOrder(node.Left, out)
	preOrder(node.Right, out)
}

//
// 解法二：手动维护栈
//
func preorderTraversalL2(root *TreeNode) []int {
	out := make([]int, 0, 10)

	tmp := root
	stack := list.New()
	stack.PushBack(tmp)

	for stack.Len() > 0 {
		tp := stack.Back()
		node := stack.Remove(tp).(*TreeNode)
		out = append(out, node.Val)

		// 由于先打印左节点，栈的特性，先压入右节点, 后打印
		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return out
}
