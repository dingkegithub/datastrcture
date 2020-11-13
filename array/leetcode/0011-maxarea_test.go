package leetcode

import "testing"

func TestMaxAreaL1(t *testing.T) {

	in := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := BucketMaxAreaL1(in)

	if area != 49 {
		t.Error("area want: 49, but: ", area)
		t.FailNow()
	}
}

func TestMaxAreaL2(t *testing.T) {

	in := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := BucketMaxAreaL2(in)

	if area != 49 {
		t.Error("area want: 49, but: ", area)
		t.FailNow()
	}
}
