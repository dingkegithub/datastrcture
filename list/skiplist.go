package list

import (
	"fmt"
	"math/rand"
)

/*

```
forwards 存储同一级上节点，就是一个子链表, forwards 数组长度代表同一级别上游多少个跳点，forwards 深度就是图上

横向 key 应该是一样的，其深度为 lvl 的链表，从当前节点跳到下一层就是:

        sn = header
		sn = sn.forward[0]    跳到第一级
        sn = sn.forward[0]    调到第二级
        sn = sn.forward[0]    调到第三级

              第一级                        第二级                                     第三级别
---------------------------          -------------------                   ---------------------------
             key
header ------lvl
             forwards+--SN ----------key, lvl, forwards+-----SN+-----------key, lvl, forwards+------SN
                     |                                 |                                     +------SN
                     |                                 |                                     +------SN
                     |                                 |
                     |                                 +-----SN+-----------key, lvl, forwards+------SN
                     |                                 |                                     +------SN
                     |                                 |                                     +------SN
                     |                                 |
                     |                                 +-----SN+-----------key, lvl, forwards+------SN
                     |                                                                       +------SN
                     |                                                                       +------SN
			         |
                     +--SN  ---------key, lvl, forwards+-----SN+-----------key, lvl, forwards+------SN
                     |                                 |                                     +------SN
                     |                                 |                                     +------SN
                     |                                 |
                     |                                 +-----SN+-----------key, lvl, forwards+------SN
                     |                                 |                                     +------SN
                     |                                 |                                     +------SN
                     |                                 +-----SN+-----------key, lvl, forwards+------SN
                     |                                                                       +------SN
                     |                                                                       +------SN
                     |
                     +--SN  ---------key, lvl, forwards+-----SN+-----------key, lvl, forwards+------SN
                                                       |                                     +------SN
                                                       |                                     +------SN
                                                       |
                                                       +-----SN+-----------key, lvl, forwards+------SN
                                                       |                                     +------SN
                                                       |                                     +------SN
                                                       +-----SN+-----------key, lvl, forwards+------SN
                                                                                             +------SN
                                                                                             +------SN


第三级 forwards = [6]
第二级 forwards = [6, 25]
第一级 forwards = [6, 9, 25]
第零级 forwards = [6, 9, 25]

3            6 ------------------------------------------------------------------->NULL
             |
            \|/
2            6 -------------------------------------------> 25------------>NULL
             |                                              |
            \|/                                            \|/
1----------> 6 ------------>9-----------------------------> 25------------>NULL
             |              |                               |
            \|/            \|/                             \|/
0   3 -----> 6  ---> 7 ---> 9 ---> 12  ---> 19 ---> 21 ---> 25 ---> 26 ---> NULL

```


*/
type SkipNode struct {
	key     int
	level   int
	forward []*SkipNode
}

func NewSkipNode(key, level int) *SkipNode {
	return &SkipNode{
		key:     key,
		level:   level,
		forward: make([]*SkipNode, level),
	}
}

type SkipList struct {
	maxLevel int
	p        float64
	level    int
	header   *SkipNode
}

func NewSkipList(maxLevel int, p float64) *SkipList {
	skipList := &SkipList{
		maxLevel: maxLevel,
		p:        p,
		header:   NewSkipNode(-1, maxLevel),
	}

	return skipList
}

// create random level for node
func (sl *SkipList) randomLevel() int {
	lvl := 0
	r := rand.Int()
	for float64(r) < sl.p && lvl < sl.maxLevel {
		lvl += 1
	}

	return lvl
}

func (sl *SkipList) Print() {
	for i := 0; i < sl.level; i++ {
		node := sl.header.forward[i]
		fmt.Print("Level ", i, ": ")

		for node != nil {
			fmt.Print(node.key)
			node = node.forward[i]
		}
		fmt.Println()
	}
}

/*
header = [key:-1, forwards: [Node-nil, Node-nil, Node-nil, Node-nil]
current = header
update = [Node-nil, Node-nil, Node-nil, Node-nil]
            |
            |
update = [header, header, header, header]
            |       |       |       |
			{key:-1, forwards: [Node-nil, Node-nil, Node-nil, Node-nil]}


current = current.forwards[0] --> Node-nil

rlevel = 2

level = rlevel = 2

newNode = {key:6, forwards: [Node-nil, Node-nil]}
                               |          |
							   nil       nil
update = [newNode, newNode, ]

MaxLevel = 3

3            6 ------------------------------------------------------------------->NULL
             |
            \|/
2            6 -------------------------------------------> 25------------>NULL
             |                                              |
            \|/                                            \|/
1----------> 6 ------------>9-----------------------------> 25------------>NULL
             |              |                               |
            \|/            \|/                             \|/
0   3 -----> 6  ---> 7 ---> 9 ---> 12  ---> 19 ---> 21 ---> 25 ---> 26 ---> NULL



3            6 ------------------------------------------------------------------->NULL
             |
            \|/
2            6 --------------------------------------------------> 25------------>NULL
             |                                                     |
            \|/                                                   \|/
1----------> 6 ------------>9------------->17--------------------> 25------------>NULL
             |              |              |                       |
            \|/            \|/            \|/                     \|/
0   3 -----> 6  ---> 7 ---> 9 ---> 12 ---> 17 ---> 19 ---> 21 ---> 25 ---> 26 ---> NULL
*/
func (sl *SkipList) Insert(key int) {
	cur := sl.header

	update := make([]*SkipNode, sl.maxLevel+1)

	// start from highest level of skip list
	// move the current reference forward while key
	// is greater than key of node next to current
	// Otherwise inserted current in update and
	// move one level down and continue search
	for i := sl.level; i >= 0; i-- {
		// 向下一层
		for cur.forward[i] != nil && cur.forward[i].key < key {
			cur = cur.forward[i]
		}

		// 记录每一层中的节点要更新的
		update[i] = cur
	}

	cur = cur.forward[0]

	if cur == nil || cur.key != key {
		rLevel := sl.randomLevel()

		if rLevel > sl.level {
			for i := sl.level + 1; i < rLevel+1; i++ {
				update[i] = sl.header
			}
			sl.level = rLevel
		}
		n := NewSkipNode(key, rLevel)

		for i := 0; i <= rLevel; i++ {
			n.forward[i] = update[i].forward[i]
			update[i].forward[i] = n
		}
	}
}

func (sl *SkipList) Search(key int) {
	cur := sl.header

	for i := sl.level; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].key < key {
			cur = cur.forward[i]
		}
	}

	cur = cur.forward[0]
	if cur != nil && cur.key == key {
		fmt.Println("Found key: ", key)
	}
}

func (sl *SkipList) Delete(key int) {
	cur := sl.header

	update := make([]*SkipNode, sl.maxLevel+1)

	for i := sl.level; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].key < key {
			cur = cur.forward[i]
		}
		update[i] = cur
	}

	cur = cur.forward[0]
	if cur != nil && cur.key == key {
		for i := 0; i <= sl.level; i++ {
			if update[i].forward[i] != cur {
				break
			}

			update[i].forward[i] = cur.forward[i]
		}

		for sl.level > 0 && sl.header.forward[sl.level] == nil {
			sl.level -= 1
		}

		fmt.Println("Success delete key: ", key)
	}
}
