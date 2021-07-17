package array

import "testing"

func TestTwoSUm(t *testing.T) {
	testcase := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}

	testcaseTarget := []int{
		9,
		6,
		6,
	}

	expected := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}

	for i := 0; i < len(testcase); i++ {
		var result []int = TwoSum(testcase[i], testcaseTarget[i])
		for j, val := range result {
			if val != expected[i][j] {
				t.Errorf("Incorrect result : expected %d got %d at test %d index %d", expected[i][j], val, testcase[i], i)
			}
		}
	}

}
