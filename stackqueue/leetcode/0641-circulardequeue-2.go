package leetcode

import (
	"sync"
)

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

   front rear
      |    |
     \+/--\+/------+------+------+------+------+------+
     |  0   |  1   |   2  |   3  |  4   |  5   |  6   |
     +------+------+------+------+------+------+------+

上一个实现本质来说不是循环队列，而是一个双端队列，设计两个指针 front, rear，

front 一直指向队列头，rear 一直指向队列的尾部，那么可以得出：

  队列空： front == rear
  队列满： (rear + 1) % capacity == front, front 和rear之间差1，当rear从队尾追
           上front 时表示满了，通俗点就是 rear 和front开始处于一个起点，然后rear
           饶了一大圈后离起点差一步时表示满，若 front = 0， 那么上图中当rear=6
           时，若再+1， （6+1）% 7 = 0，就回到了front位置，表示追上了，此时rear=6
           就是队尾，且所有元素已经填满，队列也满了

由于 rear 和 front 一个指向头，一个指向为，它们走的方向肯定是相反的，尾部插入元素

向右走，那么头部插入元素必须向左走，就是说 rear 向 (rear + 1) % capacity 的方向走，

front 向 (front - 1 + capacity) % capacity 的方向走;

同样删除元素时，front 向右走 (front + 1) % capacity, rear 向左走，就是 (rear-1 + capacity) % capacity

方向走

为什么开辟的容量要比实际大一个呢？


   front rear
      |    |
     \+/--\+/------+------+
     |  0   |  1   |   2  |
     +------+------+------+

假设容量为k = 3, 好，我们执行
     InsertFront --> true     rear=0, font=2
     InsertFront --> true     rear=0, front=1
     InsertFront --> false    reqr=0,  (rear+1) % capacity == 1 == front 队列满了
见到了吧，容量为3，我们插入了2个元素后容量被判定满，因此容量必须 k+1, 因此记住，

循环队里容量始终比给定容量k 多1，也就是实际开辟容量为 k + 1



51/51 cases passed (16 ms)
Your runtime beats 88.61 % of golang submissions
Your memory usage beats 34.4 % of golang submissions (6.7 MB)
*/
type CircularDeque struct {
	mutex    sync.RWMutex
	data     []int
	capacity int
	front    int
	rear     int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func NewCircularDeque(k int) CircularDeque {
	// 容量必须为 k + 1, 实际才能插入k个元素
	return CircularDeque{
		capacity: k + 1,
		front:    0,
		rear:     0,
		data:     make([]int, k+1, k+1),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *CircularDeque) InsertFront(value int) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.IsFull() {
		return false
	}

	// 插入头部元素，向左走，-1 的方向，注意加上capacity, 0 位置
	// 回溢出，由于循环，0 位置-1时，回到尾部，想想成一个环
	this.front = (this.front - 1 + this.capacity) % this.capacity
	this.data[this.front] = value

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *CircularDeque) InsertLast(value int) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.IsFull() {
		return false
	}

	// 这里假设了开始尾部插入的位置实际是数组的首位置，这个前面InsertFront 对应
	// 若front先插入0位置，那么首先移动 (rear + 1) % capacity 后再赋值，但是上面
	// InsertFront 先 -1 移动在插入，所以这里先赋值后插入
	this.data[this.rear] = value
	this.rear = (this.rear + 1) % this.capacity
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *CircularDeque) DeleteFront() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.IsEmpty() {
		return false
	}

	// 删除头，相当于front向队尾的方向移动，就是 +1 了
	this.front = (this.front + 1) % this.capacity
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *CircularDeque) DeleteLast() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.IsEmpty() {
		return false
	}

	// 删除尾部，相当于 rear 向 front 方向移动了，所以和 front 的移动算法一致
	this.rear = (this.rear - 1 + this.capacity) % this.capacity
	return true

}

/** Get the front item from the deque. */
func (this *CircularDeque) GetFront() int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	if this.IsEmpty() {
		return -1
	}

	// front 是移动到指定位置再放，所以直接取
	return this.data[this.front]
}

/** Get the last item from the deque. */
func (this *CircularDeque) GetRear() int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	if this.IsEmpty() {
		return -1
	}

	// rear 是先放，然后移动位置，实际位置并没有存放元素
	// 所以获取是先减后返回
	return this.data[(this.rear-1+this.capacity)%this.capacity]
}

/** Checks whether the circular deque is empty or not. */
func (this *CircularDeque) IsEmpty() bool {
	// 头尾重合，为空
	return this.front == this.rear
}

/** Checks whether the circular deque is full or not. */
func (this *CircularDeque) IsFull() bool {
	// 头尾相差1，表示满
	return (this.rear+1)%this.capacity == this.front
}
