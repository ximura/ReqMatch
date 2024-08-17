package sorted

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSortedArray(t *testing.T) {
	s := NewArray(3, func(a, b int) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}

		return -1
	}, nil)

	s.Insert(10)
	s.Insert(2)
	s.Insert(5)
	s.Insert(1)
	s.Insert(20)

	assert.Equal(t, s.Len(), 3)
	assert.DeepEqual(t, s.innerArray, []int{1, 2, 5})
}

func BenchmarkArray(b *testing.B) {
	s := NewArray(3, func(a, b int) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}

		return -1
	}, nil)

	for i := 0; i < b.N; i++ {
		s.Insert(10)
		s.Insert(2)
		s.Insert(5)
		s.Insert(1)
		s.Insert(20)

		s.Reset()
	}
}
