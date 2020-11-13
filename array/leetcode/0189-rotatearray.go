package leetcode

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

```
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
```

示例 2:

```
输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]


解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
```

说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。
*/

// 暴力解法:
//
//         缓存最右边的元素，然后逐次移动，移动完成，初始位置
// 赋值缓存值
//
// Accepted
// 35/35 cases passed (96 ms)
// Your runtime beats 13.82 % of golang submissions
// Your memory usage beats 33.42 % of golang submissions (3.2 MB)
func RotateArrayL1(nums []int, k int) {
	if len(nums) <= 0 {
		return
	}

	size := len(nums)
	cache := 0

	for i := 0; i < k; i++ {
		cache = nums[size-1]
		for j := size - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = cache
	}
}

/*


nums: 1       2        3         4      5
K: 2

counts = 0
start = 0
cur = 0
cache = 1
nums: 1       2        3         4      5
-----------------------------------------------------
       cur = (start + k) % 2  ==>  (0 + 2) % 5 = 2
	   cache = 3, nums = 1  2  1  4  5
	   counts += 1 = 1

counts = 1
start = 0
cur = 2
cache = 3
nums: 1  2  1  4  5
-------------------------------------------------------
start ≠ cur
      cur = (start + k) % 2  ==> (2 + 2) % 5 = 4
	  cache = 5, nums = 1 2 1 4 3
	  counts += 1 = 2


counts = 2
start = 0
cur = 4
cache = 5
-------------------------------------
start ≠ cur
      cur = (start + k) % 2  ==>(4 + 2) % 5 = 1
	  cache = 2, nums = 1 5 1 4 3
	  counts += 1 = 3

counts = 3
start = 0
cur = 1
cache = 2
-------------------------------------
start ≠ cur
      cur = (start + k) % 2  ==> (1 + 2) % 5 = 3
	  cache = 4, nums = 1 5 1 2 3
	  counts += 1 = 4

counts = 4
start = 0
cur = 3
cache = 4
-------------------------------------
start ≠ cur
      cur = (start + k) % 2  ==> (3 + 2) % 5 = 0
	  cache = 1, nums = 4 5 1 2 3
	  counts += 1 = 5


counts = 5
start = 0
cur = 0
cache = 1
-----------------------------------------
start = cur  =====> 内层循环结束
----------
counts = 5 == 数组长度5， 外层循环也不满足，结束


上面的推演没有遇到回到起始位置起始位置的情况，下面演示 start ≠ cur 作用，

从 0 位置的1开始经过 m1 -> m2 -> m3 -> m1 回到了原点，此时你若再继续按照

上面算法移动，没有 start ≠ cur 的限制将导致移动 6 次之后，2，4，6 不会移动

因此当移动到回到原点的情况时 start == cur 时，移动start 到下一个位置，然后

在移动

那么为什么 移动次数 counts 会限制为数组长度呢，想想也明白，因为每个元素值移动

一个位置
```
   +---------------------------m3-----------------------+
   |                                                    |
   |  +-----------m1-------+  +----m2----------------+  |
  \|/ |                   \|/ |                     \|/ |
  +----+   +-----+       +-----+      +-----+      +-----+     +-----+
  | 1  |   |  2  |       |  3  |      |  4  |      |  5  |     |  6  |
  +----+   +-----+       +-----+      +-----+      +-----+     +-----+
          /|\ |                       /|\ |                     /|\ |
           |  |                        |  |                      |  |
           |  +----------m4------------+  +-------m5-------------+  |
           |                                                        |
           |                                                        |
           |                                                        |
           +-----------------------------------------m6-------------+

```


*/
func RotateArrayL2(nums []int, k int) {
	size := len(nums)

	if size <= 0 {
		return
	}

	// counts 标记移动的次数
	counts := 0

	// start 记录每一轮起始移动位置
	// counts < size 表示每个元素只移动一次
	for start := 0; counts < size; start++ {

		// 从起始位置开始
		cur := start

		// 缓存当前值
		cache := nums[cur]

		for {
			// 计算下一个位置
			cur = (cur + k) % size

			// 缓存下一个位置的值，同时将缓存的值放到当前位置
			cache, nums[cur] = nums[cur], cache

			// 移动一个，次数+1
			counts += 1

			// 重合，本轮以start为开始的移动结束
			if cur == start {
				break
			}
		}
	}
}
