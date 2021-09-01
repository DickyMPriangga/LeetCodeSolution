package array

import "testing"

func TestRemoveElement(t *testing.T) {
	testcase := [][]int{
		{3, 2, 2, 3},
		{0, 1, 2, 2, 3, 0, 4, 2},
	}

	testVal := []int{
		3,
		2,
	}

	expected := [][]int{
		{2, 2},
		{0, 1, 4, 0, 3},
	}

	for i := 0; i < len(testcase); i++ {
		k := RemoveElement(testcase[i], testVal[i])
		if k != len(expected[i]) {
			t.Errorf("Incorrect result : expected k = %d got %d at test %d index %d", len(expected[i]), k, testcase[i], i)
		}

		for j := 0; j < k; j++ {
			if testcase[i][j] != expected[i][j] {
				t.Errorf("Incorrect result : expected %d got %d at test %d index %d", expected[i], testcase[i], testcase[i], i)
			}
		}
	}
}
