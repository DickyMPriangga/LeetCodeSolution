package linkedList

import (
	"testing"

	lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testcase := [][]int{
		{1, 2, 3, 4, 5},
		{1},
		{1, 2},
	}

	testcaseTarget := []int{
		2,
		1,
		1,
	}

	expected := [][]int{
		{1, 2, 3, 5},
		{},
		{1},
	}

	for i := 0; i < len(testcase); i++ {
		head := lList.ListFromArray(testcase[i])
		result := RemoveNthFromEnd(&head, testcaseTarget[i]).ListToArray()

		for j := 0; j < len(expected[i]); j++ {
			if expected[i][j] != result[j] {
				t.Errorf("Incorrect result : expected %d got %d at test index %d", expected[i], result, i)
				break
			}
		}
	}
}
