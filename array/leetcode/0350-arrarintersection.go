package leetcode

/*
给定两个数组，编写一个函数来计算它们的交集。



- 示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]

- 示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]


说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/

/*
时间复杂度：O(m+n)O(m+n)，其中 mm 和 nn 分别是两个数组的长度。需要遍历两个数组并对哈希表进行操作，哈希表操作的时间复杂度是 O(1)O(1)，因此总时间复杂度与两个数组的长度和呈线性关系。

空间复杂度：O(\min(m,n))O(min(m,n))，其中 mm 和 nn 分别是两个数组的长度。对较短的数组进行哈希表的操作，哈希表的大小不会超过较短的数组的长度。为返回值创建一个数组 intersection，其长度为较短的数组的长度。
*/
func IntersectL1(nums1, nums2 []int) []int {
	m := make(map[int]int)
	result := make([]int, 0, 10)
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			if m[v] > 0 {
				result = append(result, v)
				m[v] -= 1
			}
		}
	}

	return result
}
