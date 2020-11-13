package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
// 解法1： 左-根-右
//
// 递归使用了系统栈
func inorderTraversal(root *TreeNode) []int {
	out := make([]int, 0, 10)
	inorder(root, &out)
	return out
}

func inorder(node *TreeNode, out *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, out)
	*out = append(*out, node.Val)
	inorder(node.Right, out)
}

// 解法2： 手动维护栈
//
func inorderTraversalL2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	tmp := root
	out := make([]int, 10)
	stack := list.New()

	for tmp != nil || stack.Len() > 0 {
		// 左-根-右
		for tmp != nil {
			// 左边的节点全部压到栈中, 这部分相当于完成的是递归中的前两部，因为根节点的打印是在左节点打印完后，因此这里先压如跟
			// inorder(node.Left, out)
			// *out = append(*out, node.Val)
			stack.PushBack(tmp)
			tmp = tmp.Left
		}

		// 打印节点
		top := stack.Back()
		tmp = stack.Remove(top).(*TreeNode)
		out = append(out, tmp.Val)

		// 遍历右节点
		tmp = tmp.Right
	}

	return out
}
