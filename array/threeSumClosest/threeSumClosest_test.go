package array

import "testing"

func TestThreeSumClosest(t *testing.T) {
	testcase := [][]int{
		{-1, 2, 1, -4},
	}
	testcaseTarget := []int{
		1,
	}

	expected := []int{
		2,
	}

	for i := 0; i < len(testcase); i++ {
		result := ThreeSumClosest(testcase[i], testcaseTarget[i])
		if expected[i] != result {
			t.Errorf("Incorrect result : expected %d got %d at test %d index %d", expected[i], result, testcase[i], i)
		}
	}
}
