package array

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	testcase := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{},
		{0},
	}

	expected := [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
		{},
		{},
	}

	for i := 0; i < len(testcase); i++ {
		result := ThreeSum(testcase[i])
		if !compare(result, expected[i]) {
			t.Errorf("Incorrect result : expected %d got %d at test %d index %d", expected[i], result, testcase[i], i)
		}
	}
}

func compare(arr1 [][]int, arr2 [][]int) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}

	return true
}
