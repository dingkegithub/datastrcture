package leetcode

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找
出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

```
  /|\
   |
 8 +           |                             |
   |           |                             |
 7 +           |.............................|...........|
   |           |                             |           |
 6 +           |     |                       |           |
   |           |     |                       |           |
 5 +           |     |           |           |           |
   |           |     |           |           |           |
 4 +           |     |           |     |     |           |
   |           |     |           |     |     |           |
 3 +           |     |           |     |     |     |     |
   |           |     |           |     |     |     |     |
 2 +           |     |     |     |     |     |     |     |
   |           |     |     |     |     |     |     |     |
 1 +     |     |     |     |     |     |     |     |     |
   |     |     |     |     |     |     |     |     |     |
 0 +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+-->
         0     1     2     3     4     5     6     7     8     9     10
```

说白了就是那两条线组成的正方形面积最大
*/

/*
双重循环挨个遍历

Accepted
56/56 cases passed (756 ms)
Your runtime beats 5.07 % of golang submissions
Your memory usage beats 29.79 % of golang submissions (6.3 MB)
*/
func BucketMaxAreaL1(height []int) int {
	area := 0

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			w := j - i
			h := minInt(height[i], height[j])
			if a := w * h; a > area {
				area = a
			}
		}
	}
	return area
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

//
// 假设横坐标表示x 表示数组下标，纵坐标y表示下标x对应的值，
// 那么面积：
//           area = min(y<l>, y<r>) * (x<r> - x<l>)
//
//   其中r表示坐标轴上在l的右边，l表示在r的左边, 我们将横
// 坐标的距离记作 t, 那么有：
//
//           area = min(y<l>, y<r>) * t
//
// 思路是利用双指针移动，即一个指针指向数组的0下标位置，一个
// 指向数组最右边下标的位置。
//
// 如何移动：
// 假设 y<l> 小于 y<r>，那首先移动谁呢？
//
// 从公式中可以看出，面积是否最大主要因素还是 y 值较小的那个，
// 而 t 向中间移动时，怎么都会变小，也就是说决定面积的一个因素
// 可能会使面积越来越小；那么最后面积就看 y 了，移动大的，导致
// 另一个因素高度也越来越小，最终可能得到最小面积，因此只能移动
// y 值较小的那个指针
//
// Accepted
// 56/56 cases passed (20 ms)
// Your runtime beats 44.54 % of golang submissions
// Your memory usage beats 16.51 % of golang submissions (6.3 MB)
//
// 时间复杂度：O(n)
// 空间复杂度: O(1)
func BucketMaxAreaL2(height []int) int {
	area := 0
	i, j := 0, len(height)-1

	for j > i {
		a := (j - i) * minInt(height[i], height[j])
		if a > area {
			area = a
		}

		// 移动较小的指针
		if height[i] > height[j] {
			j -= 1
		} else {
			i += 1
		}
	}

	return area
}

func BucketMinAreaL2(height []int) int {
	area := 0
	i, j := 0, len(height)-1

	for j > i {
		a := (j - i) * minInt(height[i], height[j])
		if a > area {
			area = a
		}

		// 移动较小的指针
		if height[i] < height[j] {
			j -= 1
		} else {
			i += 1
		}
	}

	return area
}
