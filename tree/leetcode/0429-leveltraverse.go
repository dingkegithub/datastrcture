package leetcode

import "container/list"

/*

Tags
Companies
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。

例如，给定一个 3叉树 :

```
                         1
					   / | \
					  /  |  \
					 /   |   \
                    3    2    4
                   / \
                  /   \
                 5    6

```

输出

```
[
	[1],
	[3, 2, 4],
	[5, 6]
]
```



*/

type NtNode struct {
	Val      int
	Children []*NtNode
}

// 解法1：广度优先是一层一层遍历，是先进先出，
// 所以用队列，本地难点不再广度用队列，难点
// 在于怎么样保存每一层，看下面代码精妙之处
func NtreeLevelTreeL1(root *NtNode) [][]int {
	var res [][]int

	tmp := root
	queue := list.New()

	// 跟节点压如队列
	queue.PushBack(tmp)

	// 遍历结束条件是队列为空
	for queue.Len() > 0 {

		var level []int

		// 精妙的地方在先缓存队列当前的长度
		// 而这个长度就是我们每一层压入元素
		// 的个数
		qsize := queue.Len()

		// 切记不能省略缓存qsize，若写成 for i := 0; i < queue.Len(); i++
		// 则遍历结果不正确，会把下层的弄进来
		for i := 0; i < qsize; i++ {
			// 从队列里把整个一层的全部取出来，放到输出结果中
			item := queue.Back()
			node := queue.Remove(item).(*NtNode)
			level = append(level, node.Val)

			// 同时下一层压如队列中，由于缓存了队列长度，新加部分
			// 不会在本次中取出, 下次才会取出
			for _, v := range node.Children {
				queue.PushBack(v)
			}
		}

		// 一层遍历完, 则放入结果中
		res = append(level)
	}

	return res
}
