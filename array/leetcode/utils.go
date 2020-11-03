package leetcode

import (
	"fmt"
)

var (
	sample = [][]int{
		{0, 1},
		{1, 0},
		{1, 2},
		{0, 0, 1, 3, 0, 1},
		{1, 0, 1, 3, 0, 1},
	}

	moveZeroWant = [][]int{
		{1, 0},
		{1, 0},
		{1, 2},
		{1, 3, 1, 0, 0, 0},
		{1, 1, 3, 1, 0, 0},
	}

	//            1  2  3  4  5  6  7   8   9   10  11   12  13   14   15   16   17    18    19    20    21     22     23     24
	fiblq = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368}
	climb = []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368}
)

func PrintArrays(name string, nums []int) {
	fmt.Print(name, " : ")
	for _, v := range nums {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}

func Equal(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		fmt.Println("Error: len not equal")
		return false
	}

	for i, v := range nums1 {
		if v != nums2[i] {
			fmt.Println(v, "!=", nums2[i])
			return false
		}
	}

	return true
}
