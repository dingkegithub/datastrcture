package leetcode

import "testing"

func TestMergeTwoArray(t *testing.T) {
	num1 := make([]int, 7)
	num1[0] = 1
	num1[1] = 2
	num1[2] = 5
	num1[3] = 9

	num2 := []int{0, 1, 8}

	MergeTwoArrayL2(num1, 4, num2, 3)
	want := []int{0, 1, 1, 2, 5, 8, 9}
	for i, v := range num1 {
		if want[i] != v {
			t.Error("Merge two array want: ", i, "= ", want[i], " but: ", v)
			t.FailNow()
		}
	}
}
