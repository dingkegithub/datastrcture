package stackqueue

import (
	"github.com/dingkegithub/datastrcture/array"
	"github.com/dingkegithub/datastrcture/utils"
)

func WithQueueCapacity(c int) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		dq := target.(*defaultQueue)
		dq.capacity = c
	}
}

func WithQueueCompare(c utils.Compare) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		dq := target.(*defaultQueue)
		dq.compare = c
	}
}

type defaultQueue struct {
	capacity int
	compare  utils.Compare
	data     array.Array
}

func NewQueue(opts ...utils.ParamOptionFunc) Queue {
	q := &defaultQueue{
		capacity: 10,
		compare:  utils.CompareNum,
	}

	q.Apply(opts...)

	q.data = array.NewArrayList(
		array.WithArrayCap(q.capacity),
		array.WithArrayCompare(q.compare),
	)

	return q
}

func (queue *defaultQueue) Apply(fs ...utils.ParamOptionFunc) {
	for _, f := range fs {
		f(queue)
	}
}

// Inserts the specified element into this queue if it is possible to do so immediately
// without violating capacity restrictions, returning true upon success and error if no
// space is currently available.
func (queue *defaultQueue) Add(e interface{}) (bool, error) {
	if queue.data.Size() >= queue.capacity {
		return false, ErrBlank
	}

	return queue.data.Append(e), nil
}

// Inserts the specified element into this queue if it is possible to do so immediately
// without violating capacity restrictions.
func (queue *defaultQueue) Offer(e interface{}) bool {
	if queue.data.Size() >= queue.capacity {
		return false
	}

	return queue.data.Append(e)
}

// Retrieves and removes the head of this queue.
func (queue *defaultQueue) Remove() (interface{}, error) {
	if queue.data.Empty() {
		return nil, ErrBlank
	}

	return queue.data.RemoveIdx(0)
}

// Retrieves and removes the head of this queue, or returns nil if this queue is empty.
func (queue *defaultQueue) Poll() interface{} {
	if queue.data.Empty() {
		return nil
	}

	d, _ := queue.data.RemoveIdx(0)
	return d
}

// Retrieves, but does not remove, the head of this queue
func (queue *defaultQueue) Element() (interface{}, error) {
	if queue.data.Empty() {
		return nil, ErrBlank
	}

	return queue.data.Get(0)
}

// Retrieves, but does not remove, the head of this queue, or returns nil if this queue is
// empty
func (queue *defaultQueue) Peek() interface{} {
	if queue.data.Empty() {
		return nil
	}

	d, _ := queue.data.Get(0)

	return d
}

// queue size
func (queue *defaultQueue) Size() int {
	return queue.data.Size()
}
