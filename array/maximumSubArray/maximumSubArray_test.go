package array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubArray(t *testing.T) {
	testcase := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
		{
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
	}

	for i, tc := range testcase {
		assert.Equal(t, tc.expected, MaximumSubArray(tc.nums), "Testcase "+strconv.Itoa(i))
	}
}
