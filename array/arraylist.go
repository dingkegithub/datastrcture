package array

import (
	"fmt"
	"sync"

	"github.com/dingkegithub/datastrcture/utils"
)

func WithArrayCap(c int) utils.ParamOptionFunc {
	return func(o utils.ParamOption) {
		opt := o.(*ArrayList)
		opt.capacity = c
	}
}

func WithArrayCompare(c utils.Compare) utils.ParamOptionFunc {
	return func(o utils.ParamOption) {
		opt := o.(*ArrayList)
		opt.compare = c
	}
}

type ArrayList struct {
	mutex    sync.RWMutex
	capacity int
	size     int
	data     []interface{}
	compare  utils.Compare
}

func NewArrayList(opts ...utils.ParamOptionFunc) Array {
	arrayList := &ArrayList{
		capacity: 10,
		size:     0,
		compare:  utils.CompareNum,
	}
	arrayList.Apply(opts...)

	if arrayList.capacity <= 0 {
		arrayList.capacity = 10
	}

	arrayList.data = make([]interface{}, arrayList.capacity)

	return arrayList
}

func (a *ArrayList) Apply(opts ...utils.ParamOptionFunc) {
	for _, f := range opts {
		f(a)
	}
}

func (a *ArrayList) sureCap() {
	if a.size >= a.capacity {
		a.capacity += 10
		data := make([]interface{}, a.capacity)
		for i := 0; i < a.size; i++ {
			data[i] = a.data[i]
		}
		a.data = data
	}
}

// Inserts the specified element at the specified position in this list
func (a *ArrayList) AddIdx(idx int, e interface{}) (bool, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if idx > a.size {
		return false, ErrOutOfMemory
	}

	a.sureCap()
	for i := a.size; i > idx; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[idx] = e
	a.size += 1
	return true, nil
}

// append the element at the end of position in this list
func (a *ArrayList) Add(e interface{}) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.sureCap()
	a.data[a.size] = e
	a.size += 1
	return true
}

// Returns the element at the specified position in this list.
func (a *ArrayList) Get(idx int) (interface{}, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	if idx >= a.size {
		return false, ErrOutOfMemory
	}

	return a.data[idx], nil
}

// Returns a view of the portion of this list between the specified from, inclusive, and to, exclusive.
func (a *ArrayList) Sub(from int, to int) (Array, error) {
	if to >= a.size || from < 0 {
		return nil, ErrOutOfMemory
	}

	newArr := NewArrayList(WithArrayCap(to - from + 1))

	a.mutex.RLock()
	defer a.mutex.RUnlock()

	for i := from; i < to; i++ {
		newArr.Add(a.data[i])
	}

	return newArr, nil
}

// Returns true if this list contains no elements.
func (a *ArrayList) Empty() bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.size == 0
}

// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (a *ArrayList) Index(e interface{}) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Exception type: ", r)
			ret = -1
		}
	}()

	a.mutex.RLock()
	defer a.mutex.RUnlock()

	for i := 0; i < a.size; i++ {
		if a.compare(e, a.data[i]) == 0 {
			return i
		}
	}

	return -1
}

// Replaces the element at the specified position in this list with the specified element.
func (a *ArrayList) Set(idx int, e interface{}) (bool, error) {
	if idx < 0 || idx >= a.size {
		return false, ErrOutOfMemory
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.data[idx] = e
	return true, nil
}

// Removes the element at the specified position in this list.
func (a *ArrayList) RemoveIdx(idx int) (interface{}, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if idx < 0 || idx >= a.size {
		return false, ErrOutOfMemory
	}

	var i int
	d := a.data[idx]

	for i = idx; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.data[i] = nil
	a.size -= 1
	return d, nil
}

// Removes the first occurrence of the specified element from this list, if it is present
func (a *ArrayList) Remove(e interface{}) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("remove panic: ", r)
			ok = false
			err = ErrInvalidDataType
		}
	}()

	a.mutex.Lock()
	defer a.mutex.Unlock()

	for i := 0; i < a.size; i++ {
		if a.compare(a.data[i], e) == 0 {
			var j int
			for j = i; j < a.size-1; j++ {
				a.data[j] = a.data[j+1]
			}
			a.data[j] = nil
			a.size -= 1
			break
		}
	}

	return true, nil
}

// Appends the specified element to the end of this list.
func (a *ArrayList) Append(e interface{}) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.sureCap()
	a.data[a.size] = e
	a.size += 1
	return true
}

// Removes all of the elements from this list.
func (a *ArrayList) Clear() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	for i := range a.data {
		a.data[i] = nil
	}
	a.size = 0
}

// return element size in list
func (a *ArrayList) Size() int {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.size
}

// return total capacity
func (a *ArrayList) Capacity() int {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.capacity
}

// set capacity = size
func (a *ArrayList) Trim() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	d := make([]interface{}, a.size)
	for i := 0; i < a.size; i++ {
		d[i] = a.data[i]
	}
	a.data = d
	a.capacity = a.size
}
