package leetcode

import "testing"

func TestMaxWindow(t *testing.T) {

	sample := []int{1, 3, -1, -3, 5, 3, 6, 7, 3}

	maxWindown := MaxSlidingWindowL2(sample, 3)

	want := []int{3, 3, 5, 5, 6, 7, 7}

	for i, v := range maxWindown {
		if v != want[i] {
			t.Error("want: ", want[i], " but: ", v, i)
			t.FailNow()
		}
	}
}

func TestMaxWindowL3(t *testing.T) {

	sample := []int{1, 3, -1, -3, 5, 3, 6, 7, 3}

	maxWindown := MaxSlidingWindowL3(sample, 3)

	want := []int{3, 3, 5, 5, 6, 7, 7}

	for i, v := range maxWindown {
		if v != want[i] {
			t.Error("want: ", want[i], " but: ", v, i)
			t.FailNow()
		}
	}
}
