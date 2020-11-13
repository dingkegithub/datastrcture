package leetcode

import (
	"testing"

	"github.com/dingkegithub/datastrcture/list"
)

func TestSwapListNearPairL1(t *testing.T) {
	head := &list.ListNode{
		Data: 1,
		Next: &list.ListNode{
			Data: 2,
			Next: &list.ListNode{
				Data: 3,
				Next: &list.ListNode{
					Data: 4,
					Next: &list.ListNode{
						Data: 5,
						Next: nil,
					},
				},
			},
		},
	}

	newHead := SwapListNearPairL1(head)

	i := 0
	want := []int{2, 1, 4, 3, 5}

	for newHead != nil {
		if newHead.Data.(int) != want[i] {
			t.Error("want: ", i, want[i], " but: ", newHead.Data.(int))
			t.FailNow()
		}
		i += 1
		newHead = newHead.Next
	}
}

func TestSwapListNearPairL2(t *testing.T) {
	head := &list.ListNode{
		Data: 1,
		Next: &list.ListNode{
			Data: 2,
			Next: &list.ListNode{
				Data: 3,
				Next: &list.ListNode{
					Data: 4,
					Next: &list.ListNode{
						Data: 5,
						Next: &list.ListNode{
							Data: 6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	newHead := SwapListNearPairL2(head)

	i := 0
	want := []int{2, 1, 4, 3, 6, 5}

	for newHead != nil {
		if newHead.Data.(int) != want[i] {
			t.Error("want: ", i, want[i], " but: ", newHead.Data.(int))
			t.FailNow()
		}
		i += 1
		newHead = newHead.Next
	}
}
