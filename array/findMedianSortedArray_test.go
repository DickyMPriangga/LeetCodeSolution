package array

import "testing"

func TestFindMedianSortedArray(t *testing.T) {
	testcase := [][2][]int{
		{{1, 3}, {2}},
		{{1, 2}, {3, 4}},
		{{0, 0}, {0, 0}},
		{{}, {1}},
	}

	expected := []float64{
		2.0,
		2.5,
		0,
		1,
	}

	for i := 0; i < len(testcase); i++ {
		result := FindMedianSortedArrays(testcase[i][0], testcase[i][1])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %f got %f at test %d and %d index %d", expected[i], result, testcase[i][0], testcase[i][1], i)
		}
	}
}
