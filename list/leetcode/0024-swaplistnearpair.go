package leetcode

import "github.com/dingkegithub/datastrcture/list"

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

```
        1 --> 2 --> 3 --> 4 --> 5
                  |
                  |
                 \|/
        2 --> 1 --> 4 --> 3 --> 5
```

输入：head = [1,2,3,4]
输出：[2,1,4,3]

输入：head = []
输出：[]

输入：head = [1]
输出：[1]
*/

//
// 递归方式
//
func SwapListNearPairL1(head *list.ListNode) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = SwapListNearPairL1(next.Next)
	next.Next = head
	return next
}

//
// 非递归方式
//
func SwapListNearPairL2(head *list.ListNode) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *list.ListNode
	cur := head

	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur

		if pre != nil {
			pre.Next = next
		} else {
			head = next
		}
		pre = cur
		cur = cur.Next
	}

	return head
}
