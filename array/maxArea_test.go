package array

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	testcase := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{4, 3, 2, 1, 4},
		{1, 2, 1},
	}

	expected := []int{
		49, 1, 16, 2,
	}

	for i := 0; i < len(testcase); i++ {
		result := MaxArea(testcase[i])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %d got %d at test %d index %d", expected[i], result, testcase[i], i)
		}
	}
}
