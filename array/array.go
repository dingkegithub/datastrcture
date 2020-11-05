package array

import "fmt"

var (
	ErrOutOfMemory     = fmt.Errorf("array operation out of memory")
	ErrInvalidDataType = fmt.Errorf("invalid data type")
)

type Array interface {
	// Inserts the specified element at the specified position in this list
	AddIdx(idx int, e interface{}) (bool, error)

	// append the element at the end of position in this list
	Add(e interface{}) bool

	// Returns the element at the specified position in this list.
	Get(idx int) (interface{}, error)

	// Returns a view of the portion of this list between the specified from, inclusive, and to, exclusive.
	Sub(from int, to int) (Array, error)

	// Returns true if this list contains no elements.
	Empty() bool

	// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
	Index(e interface{}) int

	// Replaces the element at the specified position in this list with the specified element.
	Set(idx int, e interface{}) (bool, error)

	// Removes the element at the specified position in this list.
	RemoveIdx(idx int) (interface{}, error)

	// Removes the first occurrence of the specified element from this list, if it is present
	Remove(interface{}) (bool, error)

	// Appends the specified element to the end of this list.
	Append(e interface{}) bool

	// Removes all of the elements from this list.
	Clear()

	// return element size in list
	Size() int

	// return total capacity
	Capacity() int

	// set capacity = size
	Trim()
}
