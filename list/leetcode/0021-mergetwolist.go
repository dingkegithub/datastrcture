package leetcode

/*
amazon | apple | linkedin | microsoft

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

/*
时间复杂度：O(n)
空间复杂度: O(1)

时间复杂度：O(n + m) ，其中 nn 和 mm 分别为两个链表的长度。因为每次循环迭代中，
            l1 和 l2 只有一个元素会被放进合并链表中， 因此 while 循环的次数不会超过
			两个链表的长度之和。所有其他操作的时间复杂度都是常数级别的，因此总的时间
			复杂度为 O(n+m)O(n+m)。

空间复杂度：O(1) 。我们只需要常数的空间存放若干变量。

著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
208/208 cases passed (4 ms)
Your runtime beats 58.55 % of golang submissions
Your memory usage beats 68.14 % of golang submissions (2.5 MB)
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoListL1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head

	l1P, l2P := l1, l2

	for !(l1P == nil && l2P == nil) {
		if l2P == nil || (l1P != nil && l1P.Val < l2P.Val) {
			tmp.Next = l1P
			l1P = l1P.Next
		} else {
			tmp.Next = l2P
			l2P = l2P.Next
		}
		tmp = tmp.Next
	}

	return head.Next
}

/*
208/208 cases passed (4 ms)
Your runtime beats 58.55 % of golang submissions
Your memory usage beats 5.21 % of golang submissions (2.6 MB)
*/
func MergeTwoListL2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head

	l1P, l2P := l1, l2

	for !(l1P == nil && l2P == nil) {
		tmp.Next = &ListNode{}

		if l2P == nil || (l1P != nil && l1P.Val < l2P.Val) {
			tmp.Next.Val = l1P.Val
			l1P = l1P.Next
		} else {
			tmp.Next.Val = l2P.Val
			l2P = l2P.Next
		}

		tmp = tmp.Next
	}

	return head.Next
}

/*
解法3：递归

时间复杂度：O(n + m)O(n+m)，其中 nn 和 mm 分别为两个链表的长度。因为每次调用递归都会去掉
            l1 或者 l2 的头节点（直到至少有一个链表为空），函数 mergeTwoList 至多只会递归
			调用每个节点一次。因此，时间复杂度取决于合并后的链表长度，即 O(n+m)O(n+m)。

空间复杂度：O(n + m)O(n+m)，其中 nn 和 mm 分别为两个链表的长度。递归调用 mergeTwoLists 函
            数时需要消耗栈空间，栈空间的大小取决于递归调用的深度。结束递归调用时 mergeTwoLists
			函数最多调用 n+mn+m 次，因此空间复杂度为 O(n+m)O(n+m)。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists/solution/he-bing-liang-ge-you-xu-lian-biao-by-leetcode-solu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
208/208 cases passed (4 ms)
Your runtime beats 58.55 % of golang submissions
Your memory usage beats 9.93 % of golang submissions (2.6 MB)
*/
func MergeTwoListL3(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		// l1 空链表是返回l2
		return l2
	} else if l2 == nil {
		// l2 空链表是返回l1
		return l1
	} else if l1.Val < l2.Val {
		// 此处是精妙之处
		l1.Next = MergeTwoListL3(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoListL3(l1, l2.Next)
		return l2
	}
}
