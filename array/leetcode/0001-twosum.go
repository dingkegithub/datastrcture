package leetcode

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

//
// Level 1:
//       两层循环，使用两个索引，外层求出余数，内层
// 循环判断 是否等于余数，若为余数则找到
//
// 时间复杂度：O(n^2): 双层for循环
// 空间复杂度: O(1): 两个索引
func TwoSumLevel1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		// 求余数
		tmp := target - nums[i]

		// 从 i 的后一个元素开始找
		for j := i + 1; j < len(nums); j++ {
			// 为余数则返回
			if nums[j] == tmp {
				return []int{i, j}
			}
		}
	}

	return nil
}

//     array              target
// [3, -1, 2, 9, 7]  ---> 8
//
// 开一个hashmap
// map: int   --> int
//      value --> 数组索引
//
// map = {
//   3: 0,
//  -1: 1,
//   2: 2,
//   9: 3,
//   7: 4,
//}
//
// 遍历数组，一旦 target -  array[index] 的键在 map中时，说明找到
//
func TwoSumLevel2(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		// 求余数
		preKey := target - nums[i]

		// 余数是否在map中
		if _, ok := m[preKey]; ok {
			// 在则找到
			return []int{m[preKey], i}
		}
		m[nums[i]] = i
	}

	return nil
}
