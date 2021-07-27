package array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRangeSortedArray(t *testing.T) {
	testcase := []struct {
		array    []int
		target   int
		expected []int
	}{
		{
			array:    []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			array:    []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			array:    []int{},
			target:   0,
			expected: []int{-1, -1},
		},
	}

	for i, tc := range testcase {
		assert.Equal(t, tc.expected, SearchRangeSortedArray(tc.array, tc.target), "Testcase "+strconv.Itoa(i))
	}
}
