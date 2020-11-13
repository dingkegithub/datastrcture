package leetcode

import "sync"

/*
设计实现双端队列。
你的实现需要支持以下操作：

MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了

示例：

```
MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4
```

提示：

所有值的范围为 [1, 1000]
操作次数的范围为 [1, 1000]
请不要使用内置的双端队列库。


51/51 cases passed (20 ms)
Your runtime beats 50 % of golang submissions
Your memory usage beats 45.86 % of golang submissions (6.7 MB)
*/
type MyCircularDeque struct {
	mutex    sync.RWMutex
	data     []int
	capacity int
	size     int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k,
		size:     0,
		data:     make([]int, k, k),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.size >= this.capacity {
		return false
	}

	if this.size == 0 {
		this.data[0] = value
		this.size += 1
		return true
	}

	for i := this.size; i >= 0; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[0] = value
	this.size += 1
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.size >= this.capacity {
		return false
	}

	this.data[this.size] = value
	this.size += 1

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.size == 0 {
		return false
	}

	if this.size == 1 {
		this.data[0] = 0
		this.size = 0
		return true
	}

	for i := 0; i < this.size-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.size -= 1
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.size == 0 {
		return false
	}

	this.data[this.size-1] = 0
	this.size -= 1
	return true

}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	if this.size == 0 {
		return -1
	}

	return this.data[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	if this.size == 0 {
		return -1
	}
	return this.data[this.size-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity
}
