package leetcode

import (
	"testing"
)

func TestPluseOneL1(t *testing.T) {
	sample := [][]int{
		{0, 0, 0, 1},
		{0, 0, 9, 9},
		{0, 9, 9, 9},
		{9, 9, 9, 9},
		{0, 0, 9, 2},
	}
	want := [][]int{
		{0, 0, 0, 2},
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 9, 3},
	}

	for i := range sample {
		PrintArrays("in:   ", sample[i])
		PrintArrays("want: ", want[i])
		l1 := plusOneLevel1(sample[i])
		//l2 := plusOneLevel2(sample[i])
		PrintArrays("out:  ", l1)
		//PrintArrays("out:  ", l2)

		if !Equal(l1, want[i]) { // || !Equal(l2, want[i]) {
			t.FailNow()
		}
	}
}

func TestPluseOneL2(t *testing.T) {
	sample := [][]int{
		{0, 0, 0, 1},
		{0, 0, 9, 9},
		{0, 9, 9, 9},
		{9, 9, 9, 9},
		{0, 0, 9, 2},
	}
	want := [][]int{
		{0, 0, 0, 2},
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 9, 3},
	}

	for i := range sample {
		PrintArrays("in:   ", sample[i])
		PrintArrays("want: ", want[i])
		//l1 := plusOneLevel1(sample[i])
		l2 := plusOneLevel2(sample[i])
		//PrintArrays("out:  ", l1)
		PrintArrays("out:  ", l2)

		if !Equal(l2, want[i]) { // || !Equal(l2, want[i]) {
			t.FailNow()
		}
	}
}
