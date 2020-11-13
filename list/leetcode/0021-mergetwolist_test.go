package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoListL1(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	out := MergeTwoList(l1, l2)
	for out != nil {
		fmt.Print(out.Val)
		out = out.Next
	}
	fmt.Println()
}
