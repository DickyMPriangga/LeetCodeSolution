package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRotatedSortArray(t *testing.T) {
	testcase := []struct {
		array    []int
		target   int
		expected int
	}{
		{
			array:    []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			array:    []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			array:    []int{1},
			target:   0,
			expected: -1,
		},
		{
			array:    []int{1, 3},
			target:   0,
			expected: -1,
		},
		{
			array:    []int{1, 3, 5},
			target:   3,
			expected: 1,
		},
		{
			array:    []int{1, 3, 5},
			target:   5,
			expected: 2,
		},
	}

	for _, tc := range testcase {
		assert.Equal(t, tc.expected, SearchRotatedSortArray(tc.array, tc.target))
	}
}
