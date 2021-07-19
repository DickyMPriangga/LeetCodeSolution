package array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testcase := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}

	expected := [][]int{
		{1, 2},
		{0, 1, 2, 3, 4},
	}

	for i := 0; i < len(testcase); i++ {
		k := RemoveDuplicates(testcase[i])
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
