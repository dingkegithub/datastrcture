package leetcode

import "container/list"

/*
Tags
Companies
给定一个 N 叉树，返回其节点值的前序遍历。

例如，给定一个 3叉树 :


```
                         1
                       / | \
                      /  |  \
                     /   |   \
                    3    2    4
                   / \
                  /   \
                 5    6

```

返回其前序遍历: [1,3,5,6,2,4]。



说明: 递归法很简单，你可以使用迭代法完成此题吗?


*/

//
// Ntree 前序遍历
// 二叉树是根-左-右，节点变成n个是，相当于子左到右
// 把所有节点遍历一遍
//
func NtreePretraversal(root *NtNode) []int {
	out := make([]int, 10)
	ntreePreTraversal(root, &out)
}

func ntreePreTraversal(node *NtNode, out *[]int) {
	if node == nil {
		return
	}

	*out = append(*out, node.Val)
	// 遍历子节点
	for _, v := range node.Children {
		ntreePreTraversal(v, out)
	}
}

/*

```
                         1
                       / | \
                      /  |  \
                     /   |   \
                    3    2    4
                   / \
                  /   \
                 5    6

```
*/
func preorderWithStack(root *Node) []int {
	if root == nil {
		return []int{}
	}

	out := make([]int, 0, 10)

	tmp := root
	stack := list.New()

	// 头结点先压如栈
	stack.PushBack(tmp)

	for stack.Len() > 0 {
		item := stack.Back()
		node := stack.Remove(item).(*Node)
		out = append(out, node.Val)

		// 关键在这儿
		// 先序是左边的先打印，所以反一下方向先把
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
	}

	return out
}
