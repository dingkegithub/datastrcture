package leetcode

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。


说明：

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。


示例：

输入：
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出：[1,2,2,3,5,6]

分析：
     1. 数组有序
	 2. num2 放到 nums1 中
     2. nums1 的容量 >= m + n
*/

/*
解法1： 首先想到的是暴力解法，就是归并排序，需要开辟额外的空间m

时间复杂度: O(m+n)
空间复杂度: O(m)
59/59 cases passed (0 ms)
Your runtime beats 100 % of golang submissions
Your memory usage beats 7.9 % of golang submissions (2.3 MB)
*/
func MergeTwoArrayL1(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < m+n {
		return
	}

	// 由于nus1 是目标存放位置，所以开辟一个m 大小的缓存
	// 存储nums1
	cache := make([]int, m)
	copy(cache, nums1)

	i, j, k := 0, 0, 0

	// 边界条件，只要 k 不超出 m+n的范围，保证nums1 不会越界
	for k < m+n {
		// 归并排序
		if j >= n || (i < m && cache[i] < nums2[j]) {
			nums1[k] = cache[i]
			i += 1
		} else {
			nums1[k] = nums2[j]
			j += 1
		}
		k += 1
	}

}

/*
解法2： 利用双指针，从最后开始比，然后放到对应位置，因为预留空间足够存放第二数组
        所以安全，只是合并过程任然是归并排序思路

时间复杂度: O(m+n) = O(n)
空间复杂度: O(1)
m=3, 1  2  3  0 0 0
n=3, 2  5  6
k=6

num1P = 2
num2P = 2
k = 3 + 3 - 1 = 5
--------------------------------------
     1  2  3  0  0 6

num1P = 2
num2P = 1
k = 4
--------------------------------------
     1  2  3  0  5 6

num1P = 2
num2P = 0
k = 3
--------------------------------------
     1  2  3  3  5 6

num1P = 1
num2P = 0
k = 2
--------------------------------------
     1  2  2  3  5 6

num1P = 0
num2P = -1
k = 1
--------------------------------------
     1  2  2  3  5 6

num1P = 0
num2P = -1
k = 0
--------------------------------------
     1  2  2  3  5 6

num1P = -1
num2P = -1
k = -1
结束

Accepted
59/59 cases passed (0 ms)
Your runtime beats 100 % of golang submissions
Your memory usage beats 41.6 % of golang submissions (2.3 MB)
*/
func MergeTwoArrayL2(nums1 []int, m int, nums2 []int, n int) {

	if cap(nums1) < m+n {
		return
	}

	// 双指针，分别指向两个数组最后一个元素
	num1P := m - 1
	num2P := n - 1

	// 移动次数
	k := m + n - 1

	// 从最后开始放元素, 思路还是归并，所不同是不需呀额外空间
	for i := k; i >= 0; i-- {
		if num2P < 0 || (num1P >= 0 && nums1[num1P] >= nums2[num2P]) {
			nums1[i] = nums1[num1P]
			num1P -= 1
		} else {
			nums1[i] = nums2[num2P]
			num2P -= 1
		}
	}
}
