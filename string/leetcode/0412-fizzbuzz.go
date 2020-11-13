package leetcode

import (
	"fmt"
	"strings"
)

//
// 解法0：和解法1思路一致，但是不够优雅，模操作过多
func FizzBuzzL0(n int) []string {
	if n <= 0 {
		return nil
	}

	out := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			var s strings.Builder

			if i%3 == 0 {
				s.WriteString("Fizz")
			}

			if i%5 == 0 {
				s.WriteString("Buzz")
			}
			out[i-1] = s.String()
		} else {
			out[i-1] = fmt.Sprintf("%d", i)
		}
	}
	return out
}

//
// 解法1：对 3 和 5 求余，通过是否为0来判断是否替换成Fizz 或 Buzz
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
//
func FizzBuzzL1(n int) []string {
	out := make([]string, n)

	for i := 1; i <= n; i++ {
		dived3, dived5 := false, false

		if i%3 == 0 {
			dived3 = true
		}

		if i%5 == 0 {
			dived5 = true
		}

		if dived3 && dived5 {
			out[i-1] = "FizzBuzz"
		} else if dived3 {
			out[i-1] = "Fizz"
		} else if dived5 {
			out[i-1] = "Buzz"
		} else {
			out[i-1] = fmt.Sprintf("%d", i)
		}
	}
	return out
}

//
// 解法二：比较巧妙，观察了规律，总结出fizz和buzz的规律
//
func FizzBuzzL2(n int) []string {
	if n <= 0 {
		return nil
	}

	out := make([]string, n)

	for i, fizz, buzz := 1, 0, 0; i <= n; i++ {
		fizz += 1
		buzz += 1

		if fizz == 3 && buzz == 5 {
			out[i-1] = "FizzBuzz"
		} else if fizz == 3 {
			out[i-1] = "Fizz"
		} else if buzz == 5 {
			out[i-1] = "Buzz"
		} else {
			out[i-1] = fmt.Sprintf("%d", i)
		}
	}

	return out
}

//
// 解法三：来自网友，引入hashmap 的原因是应对需求变化，比如再加个7 Jizz不用修改代码
//
// 思路很好，但是这里有一个问题，map 的遍历并非每次key 都是按照固定顺序，所以得出的结果
// 可能错误
func FizzBuzzL3(n int) []string {
	out := make([]string, n)

	numMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	for i := 1; i < n; i++ {
		var s strings.Builder

		for k, v := range numMap {
			if n%k == 0 {
				s.WriteString(v)
			}
		}

		if s.Len() <= 0 {
			out[i-1] = fmt.Sprintf("%d", i)
		} else {
			out[i-1] = s.String()
		}
	}

	return out
}
