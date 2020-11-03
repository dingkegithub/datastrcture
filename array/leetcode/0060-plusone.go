package leetcode

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

输入: [9, 9]
输出: [1, 0, 0]
解释: 输入数组表示数字 100
*/
func plusOneLevel1(digits []int) []int {
	size := len(digits)

	for i := size - 1; i >= 0; i-- {
		// 预加 1，看是否要进位
		if digits[i]+1 >= 10 {
			// 有进位，当前为置0
			digits[i] = 0
		} else {
			// 无需进位，直接+1，返回
			digits[i] += 1
			return digits
		}
	}

	out := make([]int, size+1)
	out[0] = 1

	return out
}

//
// 简洁写法
//
func plusOneLevel2(digits []int) []int {
	size := len(digits)

	for i := size - 1; i >= 0; i-- {
		// 直接进来 +1
		digits[i] += 1

		// 取10以内值设置当前位
		digits[i] %= 10

		// 若当前为不为0，表示无进位，直接返回
		if digits[i] != 0 {
			return digits
		}
	}

	// 溢出，直接多开一个长度，0 位设置1就行
	// 主要处理 [9, 9] 这种现象
	out := make([]int, size+1)
	out[0] = 1

	return out
}

func plusOneLevel3(digits []int) []int {
	size := len(digits)

	for i := size - 1; i >= 0; i-- {
		// 更好的边界条件，代码更简洁
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[0] = 0
		}
	}

	out := make([]int, size+1)
	out[0] = 1

	return out
}
