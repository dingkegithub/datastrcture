package leetcode

func MoveZero(nums []int) {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k {
				nums[i], nums[k] = nums[k], nums[i]
			}

			k += 1
		}
	}
}
