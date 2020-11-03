package leetcode

import "math"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

1: 1          -> 1 种

2: 2          -> 2 种
   1 1

3: 1 1 1      -> 3 种
   1 2
   2 1

4: 1 1 1 1     -> 5 种
   1 1 2
   1 2 1
   2 1 1
   2 2

5: 1 1 1 1 1     -> 8 种
   1 1 1 2
   1 1 2 1
   1 2 1 1
   2 1 1 1
   1 2 2
   2 2 1
   2 1 2

6： 1 1 1 1 1 1   -> 13

    1 1 1 1 2
    1 1 1 2 1
    1 1 2 1 1
    1 2 1 1 1
    2 1 1 1 1

    2 2 1 1
    2 1 2 1
    2 1 1 2
    1 1 2 2
    1 2 1 2
    1 1 2 2

    2 2 2

斐波拉契数列 f(n) = f(n-1) + f(n-2)
       0 -> 0
       1 -> 1
       2 -> 1
       3 -> 2
       4 -> 3
       5 -> 5
       6 -> 8
       7 -> 13

题目问题和斐波拉契数列切合，因此用斐波拉契解决

只要我们再问题中，当为 1 时返回 1， 2时返回2，则后续就符合斐波拉契了
*/

// 方法1：直接用斐波拉契函数
// 时间复杂度: O(2^n)
// 空间复杂度：O(1)
func ClimbStairsLevel1(n int) int {
	// 2以内没有和斐波拉契数列切合
	// 所以直接设置其值
	if n <= 2 {
		return n
	}

	// f(n) = f(n-1) + f(n-2)
	return ClimbStairsLevel1(n-1) + ClimbStairsLevel1(n-2)
}

/*
以5为例

                            6
                          /   \
                         /     \
                        /       \
                       /         \
                      /           \
                     5            ④👄
                  /   \          /   \
                 /     \        /     \
               ④ 👄    3😂    3😂     2
              /  \    /  \   /  \
             /    \  2   1  2    1
            /      \
           3😂      2
          /  \
         2   1

3 计算了3次，4计算了2次, 若把这些缓存下来操作将简化很多


                            6
                          /   \
                         /     \
                       /        \
                     /           \
                   /              \
                 5                 ④👄
               /   \
             /      \
           ④ 👄     3😂
          /  \
         /    \
        /      \
      3😂      2
    /  \
   2   1

时间复杂度： O(n)
空间复杂度:  O(n)
*/
func ClimbStairsLevel2(n int) int {
	cache := make([]int, n+1)
	return climbStairsLevel2(n, cache)
}

func climbStairsLevel2(n int, cache []int) int {
	if n <= 2 {
		return n
	}

	// 缓存里有的不计算，直接返回
	if cache[n] != 0 {
		return cache[n]
	}

	return climbStairsLevel2(n-1, cache) + climbStairsLevel2(n-2, cache)
}

//
// 非递归话优化
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func ClimbStairsLevel3(n int) int {
	if n <= 2 {
		return n
	}

	cache := make([]int, n+1)
	cache[1] = 1
	cache[2] = 2

	for i := 3; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}

//
// 因为只需要知道n-1 和 n-2的值，因此只要暂存他们就行
//
func ClimbStairsLevel4(n int) int {
	if n <= 2 {
		return n
	}

	// 暂存 f(n-2)
	prePre := 1

	// 暂存 f(n-1)
	pre := 2

	// 暂存 f(n)
	steps := 0

	for i := 3; i <= n; i++ {
		// f(n) = f(n-1) + f(n-2)
		steps = pre + prePre
		prePre = pre
		pre = steps
	}

	return steps
}

/*
通项公式:

             1       1 + √5   n            1  - √5        n
     f(n) = --- [ ( -------- )    -  ( ------------------)   ]
	        √5         2                       2

--------+---------------------------------------------------------------------
   n    |      1        2       3     4     5      6    7    8    9   10
--------+----------------------------------------------------------------------
爬楼梯  |      1        2       3     5     8      13   21   34   55  89   ...
--------+----------------------------------------------------------------------
斐波拉契|      1        1       2     3     5      8    13   21   34  55   ...
--------+----------------------------------------------------------------------

得出揭露： climb(n) = fblq(n+1)

√5  =  5^0.5

时间复杂度：O(logn)，pow 方法将会用去 O(logn) 的时间。
空间复杂度：O(1)
*/
func ClimbStairsLevel5(n int) int {
	if n < 2 {
		return n
	}

	sqrt5 := math.Pow(5., 0.5)
	a := (1. + sqrt5) / 2.
	b := (1. - sqrt5) / 2.

	c := math.Pow(a, float64(n+1)) - math.Pow(b, float64(n+1))
	return int(c / sqrt5)
}
