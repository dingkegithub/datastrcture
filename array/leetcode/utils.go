package leetcode

import "fmt"

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
)

func PrintArrays(nums []int) {
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
