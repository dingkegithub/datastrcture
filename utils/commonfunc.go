package utils

import (
	"fmt"
	"reflect"
)

//  1: a > b
//  0: a == b
// -1: a < b
type Compare func(a, b interface{}) int

func CompareNum(a, b interface{}) int {
	tp := reflect.TypeOf(a)

	switch tp.Kind() {
	case reflect.Int:
		aI := a.(int)
		bI := b.(int)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Int8:
		aI := a.(int8)
		bI := b.(int8)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Int16:
		aI := a.(int16)
		bI := b.(int16)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Int32:
		aI := a.(int32)
		bI := b.(int32)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Int64:
		aI := a.(int64)
		bI := b.(int64)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Uint:
		aI := a.(uint)
		bI := b.(uint)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Uint8:
		aI := a.(uint8)
		bI := b.(uint8)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Uint16:
		aI := a.(uint16)
		bI := b.(uint16)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Uint32:
		aI := a.(uint32)
		bI := b.(uint32)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Uint64:
		aI := a.(uint64)
		bI := b.(uint64)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Float32:
		aI := a.(float32)
		bI := b.(float32)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	case reflect.Float64:
		aI := a.(float64)
		bI := b.(float64)
		if aI > bI {
			return 1
		} else if aI == bI {
			return 0
		} else {
			return -1
		}
	default:
		panic(fmt.Sprint("invalid type", tp.Kind()))
	}
}
