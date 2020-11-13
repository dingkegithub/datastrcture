package leetcode

import "testing"

func TestDuplicatL1(t *testing.T) {
	in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	PrintArrays("in:  ", in)
	size := RemoveDuplicate(in)
	if size != 5 {
		t.Error("Non dumplicate size want: 5, but: ", size)
		t.FailNow()
	}

	PrintArrays("out: ", in)
}
