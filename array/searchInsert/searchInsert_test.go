package array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	testcase := []struct {
		array    []int
		target   int
		expected int
	}{
		{
			array:    []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			array:    []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			array:    []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			array:    []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			array:    []int{1},
			target:   0,
			expected: 0,
		},
		{
			array:    []int{1, 3},
			target:   2,
			expected: 1,
		},
		{
			array:    []int{1},
			target:   1,
			expected: 0,
		},
	}

	for i, tc := range testcase {
		assert.Equal(t, tc.expected, SearchInsert(tc.array, tc.target), "Testcase "+strconv.Itoa(i))
	}
}
