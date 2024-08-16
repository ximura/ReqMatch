package sorted

import (
	"bytes"
	"slices"
)

type Array[T any] struct {
	innerArray []T

	cmp func(a, b T) int
	str func(T) string
}

func NewArray[T any](limit int, cmp func(a, b T) int, str func(T) string) Array[T] {
	return Array[T]{
		innerArray: make([]T, 0, limit),
		cmp:        cmp,
		str:        str,
	}
}

func (s *Array[T]) Insert(v T) {
	i, _ := slices.BinarySearchFunc(s.innerArray, v, s.cmp) // find slot
	c := cap(s.innerArray)
	n := len(s.innerArray)
	if i >= c {
		// out of array capacity
		return
	}

	if n == i {
		// we stil in array capacity, so just append new element
		s.innerArray = append(s.innerArray, v)
		return
	}

	a := s.innerArray
	if n != c {
		// if we still have capacity create new element in tail
		a = s.innerArray[:n+1]
	}
	copy(a[i+1:], a[i:])
	a[i] = v
	s.innerArray = a
}

func (s *Array[T]) Reset() {
	s.innerArray = s.innerArray[:0]
}

func (s *Array[T]) Len() int {
	return len(s.innerArray)
}

func (s *Array[T]) Marshal() ([]byte, error) {
	var builder bytes.Buffer
	n := len(s.innerArray)

	// id: uint32 max 10 bytes for string representation or MaxValue
	// delay: 10 seconds is up to 1e+10 so we would need 10 bytes

	builder.Grow(n*23 + n - 1 + 2) // 23 = 10 + 10 (for id and delay) + 3 (for "" and :); n - 1  - comma character in json; 2 - for {}
	builder.WriteString("{")
	for i := range s.innerArray {
		v := s.innerArray[i]
		builder.WriteString(s.str(v))
		if i < n-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("}")

	return builder.Bytes(), nil
}
