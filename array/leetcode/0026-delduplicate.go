package leetcode

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

- 示例 1:

```
给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
```

示例 2:

```
给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
```

题目：额外空间是 O(1), 这就要求不能以hash等去解决
      数组已经排好序
*/

// 用 pre 指针标记重复的为位置，一旦重复的位置被后面不重复的元素
// 填充之后移动1 位
//
//
// 161/161 cases passed (8 ms)
// Your runtime beats 88.25 % of golang submissions
// Your memory usage beats 55.33 % of golang submissions (4.6 MB)
func RemoveDuplicate(nums []int) int {
	// 标记重复的位置
	pre := 1

	for i := 1; i < len(nums); i++ {

		// 不重复则pre位置填不重复元素
		// 指针向右移动
		if nums[i] != nums[i-1] {
			nums[pre] = nums[i]
			pre += 1
		}
	}

	return pre
}
