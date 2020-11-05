package stackqueue

import (
	"fmt"
)

var (
	ErrNoSpace = fmt.Errorf("no space")
)

type Stack interface {
	// Tests if this stack is empty.
	Empty() bool

	// Looks at the object at the top of this stack without removing it from the stack.
	Peek() interface{}

	// Removes the object at the top of this stack and returns that object as the value of this function.
	Pop() interface{}

	// Pushes an item onto the top of this stack.
	Push(e interface{}) bool

	// Returns the 1-based position where an object is on this stack.
	Search(e interface{}) int
}

type Queue interface {
	// Inserts the specified element into this queue if it is possible to do so immediately
	// without violating capacity restrictions, returning true upon success and error if no
	// space is currently available.
	Add(e interface{}) (bool, error)

	// Inserts the specified element into this queue if it is possible to do so immediately
	// without violating capacity restrictions.
	Offer(e interface{}) bool

	// Retrieves and removes the head of this queue.
	Remove() (interface{}, error)

	// Retrieves and removes the head of this queue, or returns nil if this queue is empty.
	Poll() interface{}

	// Retrieves, but does not remove, the head of this queue
	Element() (interface{}, error)

	// Retrieves, but does not remove, the head of this queue, or returns nil if this queue is
	// empty
	Peek() interface{}

	// queue size
	Size() int
}

type Dequeue interface {
	Queue

	// Inserts the specified element at the front of this deque if it is possible to do so immediately
	// without violating capacity restrictions, return an ErrNoSpace if no space is currently available.
	AddFirst(e interface{}) error

	// Inserts the specified element at the end of this deque if it is possible to do so immediately without
	// violating capacity restrictions, return an ErrNoSpace if no space is currently available.
	AddLast(e interface{}) error

	// Inserts the specified element at the front of this deque unless it would violate capacity restrictions.
	OfferFirst(e interface{}) bool

	// Inserts the specified element at the end of this deque unless it would violate capacity restrictions.
	OfferLast(e interface{}) bool

	// Retrieves and removes the first element of this deque.
	RemoveFirst() (interface{}, error)

	// Retrieves and removes the last element of this deque.
	RemoveLast() (interface{}, error)

	// Retrieves and removes the first element of this deque, or returns nil if this deque is empty
	PollFirst() interface{}

	// Retrieves and removes the last element of this deque, or returns nil if this deque is empty.
	PollLast() interface{}

	// Retrieves, but does not remove, the first element of this deque.
	GetFirst() interface{}

	// Retrieves, but does not remove, the last element of this deque.
	GetLast() interface{}

	// Retrieves, but does not remove, the first element of this deque, or returns nil if this deque is empty
	PeekFirst() interface{}

	// Retrieves, but does not remove, the last element of this deque, or returns nil if this deque is empty.
	PeekLast() interface{}
}
