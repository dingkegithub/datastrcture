package stackqueue

import (
	"github.com/dingkegithub/datastrcture/array"
	"github.com/dingkegithub/datastrcture/utils"
)

func WithCap(c int) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		stack := target.(*defaultStack)
		stack.capacity = c
	}
}

func WithCompare(c utils.Compare) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		stack := target.(*defaultStack)
		stack.compare = c
	}
}

type defaultStack struct {
	capacity int
	compare  utils.Compare
	data     array.Array
}

func NewStack(opts ...utils.ParamOptionFunc) Stack {
	stack := &defaultStack{
		capacity: 10,
		compare:  utils.CompareNum,
	}
	stack.Apply(opts...)

	stack.data = array.NewArrayList(
		array.WithArrayCap(stack.capacity),
		array.WithArrayCompare(stack.compare))

	return stack
}

func (stack *defaultStack) Apply(opts ...utils.ParamOptionFunc) {
	for _, f := range opts {
		f(stack)
	}
}

// Tests if this stack is empty.
func (stack *defaultStack) Empty() bool {
	return stack.data.Empty()
}

// Looks at the object at the top of this stack without removing it from the stack.
func (stack *defaultStack) Peek() (interface{}, error) {
	if stack.data.Empty() {
		return nil, ErrBlank
	}

	return stack.data.Get(stack.data.Size() - 1)
}

// Removes the object at the top of this stack and returns that object as the value of this function.
func (stack *defaultStack) Pop() (interface{}, error) {
	if stack.data.Empty() {
		return nil, ErrBlank
	}

	return stack.data.RemoveIdx(stack.data.Size() - 1)
}

// Pushes an item onto the top of this stack.
func (stack *defaultStack) Push(e interface{}) bool {
	return stack.data.Append(e)
}

// Returns the 1-based position where an object is on this stack.
func (stack *defaultStack) Search(e interface{}) int {
	return stack.data.Index(e)
}
