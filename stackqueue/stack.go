package stackqueue

import "github.com/dingkegithub/datastrcture/array"

type defaultStack struct {
	data array.Array
}

func NewStack() Stack {
	return &defaultStack{}
}
