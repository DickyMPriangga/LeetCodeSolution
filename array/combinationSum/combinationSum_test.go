package array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationArray(t *testing.T) {
	testcase := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			candidates: []int{1},
			target:     1,
			expected:   [][]int{{1}},
		},
		{
			candidates: []int{1},
			target:     2,
			expected:   [][]int{{1, 1}},
		},
	}

	for i, tc := range testcase {
		assert.Equal(t, tc.expected, CombinationSum(tc.candidates, tc.target), "Testcase "+strconv.Itoa(i))
	}
}
