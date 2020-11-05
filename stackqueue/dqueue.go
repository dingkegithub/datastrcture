package stackqueue

import (
	"github.com/dingkegithub/datastrcture/array"
	"github.com/dingkegithub/datastrcture/utils"
)

func WithDqueueCapacity(c int) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		dq := target.(*defaultDqueue)
		dq.capacity = c
	}
}

func WithDqueueCompare(c utils.Compare) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		dq := target.(*defaultDqueue)
		dq.compare = c
	}
}

type defaultDqueue struct {
	capacity int
	compare  utils.Compare
	data     array.Array
}

func NewDequeue(opts ...utils.ParamOptionFunc) Dequeue {
	dq := &defaultDqueue{
		capacity: 10,
		compare:  utils.CompareNum,
	}
	dq.Apply(opts...)
	dq.data = array.NewArrayList(
		array.WithArrayCap(dq.capacity),
		array.WithArrayCompare(dq.compare),
	)

	return dq
}

func (dq *defaultDqueue) Apply(fs ...utils.ParamOptionFunc) {
	for _, f := range fs {
		f(dq)
	}
}

// Inserts the specified element at the front of this deque if it is possible to do so immediately
// without violating capacity restrictions, return an ErrNoSpace if no space is currently available.
func (dq *defaultDqueue) AddFirst(e interface{}) error {
	if dq.data.Size() >= dq.capacity {
		return ErrNoSpace
	}
	_, err := dq.data.AddIdx(0, e)
	return err
}

// Inserts the specified element at the end of this deque if it is possible to do so immediately without
// violating capacity restrictions, return an ErrNoSpace if no space is currently available.
func (dq *defaultDqueue) AddLast(e interface{}) error {
	if dq.data.Size() >= dq.capacity {
		return ErrNoSpace
	}
	ok := dq.data.Append(e)
	if !ok {
		return ErrNoSpace
	}
	return nil
}

// Inserts the specified element at the front of this deque unless it would violate capacity restrictions.
func (dq *defaultDqueue) OfferFirst(e interface{}) bool {
	if dq.data.Size() >= dq.capacity {
		return false
	}
	ok, _ := dq.data.AddIdx(0, e)
	return ok
}

// Inserts the specified element at the end of this deque unless it would violate capacity restrictions.
func (dq *defaultDqueue) OfferLast(e interface{}) bool {
	if dq.data.Size() >= dq.capacity {
		return false
	}
	return dq.data.Append(e)
}

// Retrieves and removes the first element of this deque.
func (dq *defaultDqueue) RemoveFirst() (interface{}, error) {
	if dq.data.Empty() {
		return nil, ErrBlank
	}

	return dq.data.RemoveIdx(0)
}

// Retrieves and removes the last element of this deque.
func (dq *defaultDqueue) RemoveLast() (interface{}, error) {
	if dq.data.Empty() {
		return nil, ErrBlank
	}

	return dq.data.RemoveIdx(dq.data.Size() - 1)
}

// Retrieves and removes the first element of this deque, or returns nil if this deque is empty
func (dq *defaultDqueue) PollFirst() interface{} {
	if dq.data.Empty() {
		return nil
	}

	d, _ := dq.data.RemoveIdx(0)
	return d
}

// Retrieves and removes the last element of this deque, or returns nil if this deque is empty.
func (dq *defaultDqueue) PollLast() interface{} {
	if dq.data.Empty() {
		return nil
	}

	d, _ := dq.data.RemoveIdx(dq.data.Size() - 1)
	return d
}

// Retrieves, but does not remove, the first element of this deque.
func (dq *defaultDqueue) GetFirst() (interface{}, error) {
	if dq.data.Empty() {
		return nil, ErrBlank
	}

	return dq.data.Get(0)
}

// Retrieves, but does not remove, the last element of this deque.
func (dq *defaultDqueue) GetLast() (interface{}, error) {
	if dq.data.Empty() {
		return nil, ErrBlank
	}

	return dq.data.Get(dq.data.Size() - 1)
}

// Retrieves, but does not remove, the first element of this deque, or returns nil if this deque is empty
func (dq *defaultDqueue) PeekFirst() interface{} {
	if dq.data.Empty() {
		return nil
	}

	e, _ := dq.data.Get(0)
	return e
}

// Retrieves, but does not remove, the last element of this deque, or returns nil if this deque is empty.
func (dq *defaultDqueue) PeekLast() interface{} {
	if dq.data.Empty() {
		return nil
	}

	e, _ := dq.data.Get(dq.data.Size() - 1)
	return e
}
