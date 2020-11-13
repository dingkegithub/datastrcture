package leetcode

import (
	"container/list"

	"github.com/dingkegithub/datastrcture/stackqueue"
)

/*
Tags
Companies
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。



进阶：

你能在线性时间复杂度内解决此题吗？



示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/

/*
解法1：暴力解

时间复杂度: O(k * n)
空间复杂度: O(n)

leetcode 运行超时
*/
func MaxSlidingWindow(nums []int, k int) []int {

	if len(nums) <= 0 {
		return nil
	}

	windowMaxNum := make([]int, len(nums)-k+1)

	for i := 0; i <= len(nums)-k; i++ {

		maxItem := nums[i]
		for j := i + 1; j < i+k; j++ {
			if maxItem < nums[j] {
				maxItem = nums[j]
			}
		}
		windowMaxNum[i] = maxItem
	}

	return windowMaxNum
}

/*
        0  1   2   3  4  5  6  7  8
sample {4, 3,  5, -3, 5, 3, 6, 7, 3}
  0,2-> 4------5
     1,3-> 3------- 5
       2,4->   5------ 5
          3,5->   -3----- 5
              4,6->   5-----6
                 5,7->   3-----7
                    6,8->   6-----7


k       3
------------------------------------------------------------------------------------------------------

```
# i=0:

      +-----+-----+-----+
      |  4  |     |     |
      +-----+-----+-----+

# i=1:

      +-----+-----+-----+
      |  4  |  3  |     |
      +-----+-----+-----+

# i=2: 插入5，不满足单调递减，队尾开始扫描删除不满足元素

      +-----+-----+-----+
      |  4  |  3  |     |
      +-----+-----+-----+

      +-----+-----+-----+
      |  4  |     |     |       3 不满足 > 5 的条件，删除
      +-----+-----+-----+

      +-----+-----+-----+
      |     |     |     |       4 不满足 > 5 的条件，删除
      +-----+-----+-----+

      +-----+-----+-----+
      |  5  |     |     |
      +-----+-----+-----+
```
到此，第1窗口的最大值已经求出，       res = [5, ]


        0  1   2   3  4  5  6  7  8
sample {4, 3,  5, -3, 5, 3, 6, 7, 3}
k       3
```
i=3

      +-----+-----+-----+
      |  5  | -3  |     |
      +-----+-----+-----+

	 第2窗口的最大值已经求出，       res = [5, 5]

i=4, 插入5

      +-----+-----+-----+
      |  5  | -3  |     |
      +-----+-----+-----+

      +-----+-----+-----+
      |  5  |     |     |   -3 < 5 删除
      +-----+-----+-----+

	  # 插入当前元素

      +-----+-----+-----+
      |  5  |  5  |     |
      +-----+-----+-----+

	 第3窗口的最大值已经求出，       res = [5, 5, 5]

i=5, 插入3

     i - k = 5 - 3 = 2, nums[2] = 5, 等于队列的首元素，所以对首为前一个窗口的，剔除

      +-----+-----+-----+
      |  5  |     |     |
      +-----+-----+-----+

      +-----+-----+-----+
      |  5  |  3  |     |  加入当前元素
      +-----+-----+-----+

	 第4窗口的最大值已经求出，       res = [5, 5, 5, 5]



i=6, 插入6

      +-----+-----+-----+
      |  5  |  3  |     |
      +-----+-----+-----+

      +-----+-----+-----+
      |  5  |     |     |    3 < 6 删除
      +-----+-----+-----+

      +-----+-----+-----+
      |     |     |     |    5 < 6 删除
      +-----+-----+-----+

      +-----+-----+-----+
      |  6  |     |     |    插入新元素
      +-----+-----+-----+

	 第5窗口的最大值已经求出，       res = [5, 5, 5, 5, 6]


i=7, 插入7

      +-----+-----+-----+
      |  6  |     |     |
      +-----+-----+-----+

      +-----+-----+-----+
      |     |     |     |    6 < 7 删除
      +-----+-----+-----+

      +-----+-----+-----+
      |  7  |     |     |    插入新元素
      +-----+-----+-----+

	 第7窗口的最大值已经求出，       res = [5, 5, 5, 5, 6, 7]


i=8, 插入3

      +-----+-----+-----+
      |  7  |  3  |     |
      +-----+-----+-----+

	 第7窗口的最大值已经求出，       res = [5, 5, 5, 5, 6, 7, 7]
```

59/59 cases passed (284 ms)
Your runtime beats 12.94 % of golang submissions
Your memory usage beats 32.84 % of golang submissions (9.4 MB)
*/
func MaxSlidingWindowL2(nums []int, k int) []int {
	if len(nums) <= 0 {
		return nil
	}

	// 对于n个长度的数组，有 n -k + 1 个窗口
	res := make([]int, len(nums)-k+1)

	// 首先把前k个元素放到队列里面，保持头到尾单调递减，若不满足则
	// 从队列删除小的元素，保留最大的几个按递减
	dq := stackqueue.NewDequeue(stackqueue.WithDqueueCapacity(k))
	for i := 0; i < k; i++ {
		// 即将插入的元素比队尾大，则不满足单调递减，直接删除队尾
		// 遍历队列，删除小于比当前待插入元素小的元素
		for !dq.Empty() && dq.PeekLast().(int) < nums[i] {
			dq.RemoveLast()
		}

		// 插入当前元素
		dq.AddLast(nums[i])
	}

	// 当前的头肯定是 前 k 个元素的最大值，也就是第一个窗口的大小
	res[0] = dq.PeekFirst().(int)

	// 这一步要理解，从此刻开始，只要插入，就满足一个窗口
	for i := k; i < len(nums); i++ {
		// 从此刻起，i 每前进一步，都将产生一个窗口

		// 这一步很关键
		// 这一步保证窗口维持 k 大小, 若队列里的第一个元素
		// 是前一个队列的，则剔除, 加入一个意味着窗口应该为四
		// 若队列里面有前一个窗口的首元素，则删除
		if !dq.Empty() && dq.PeekFirst() == nums[i-k] {
			dq.RemoveFirst()
		}

		// 从队尾开始，遍历删除不满足队列递减条件的队尾
		for !dq.Empty() && dq.PeekLast().(int) < nums[i] {
			dq.RemoveLast()
		}

		// 插入当前元素
		dq.AddLast(nums[i])

		// 第 i-k+1 个窗口的最大值
		res[i-k+1] = dq.PeekFirst().(int)
	}

	return res
}

func MaxSlidingWindowL3(nums []int, k int) []int {
	dq := list.New()

	for i := 0; i < k; i++ {
		for dq.Len() > 0 {
			back := dq.Back()
			if back.Value.(int) < nums[i] {
				dq.Remove(back)
			} else {
				break
			}
		}
		dq.PushBack(nums[i])
	}

	res := make([]int, len(nums)-k+1)
	res[0] = dq.Front().Value.(int)

	for i := k; i < len(nums); i++ {
		if dq.Len() > 0 && dq.Front().Value.(int) == nums[i-k] {
			dq.Remove(dq.Front())
		}

		for dq.Len() > 0 && dq.Back().Value.(int) < nums[i] {
			dq.Remove(dq.Back())
		}

		dq.PushBack(nums[i])
		res[i-k+1] = dq.Front().Value.(int)
	}

	return res
}
