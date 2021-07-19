package linkedList

import (
	"testing"

	lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"
)

func TestAddTwoNumbers(t *testing.T) {
	testcase := [3][2]lList.ListNode{
		{
			lList.ListFromArray([]int{2, 4, 3}),
			lList.ListFromArray([]int{5, 6, 4}),
		},
		{
			lList.ListFromArray([]int{0}),
			lList.ListFromArray([]int{0}),
		},
		{
			lList.ListFromArray([]int{9, 9, 9, 9, 9, 9, 9}),
			lList.ListFromArray([]int{9, 9, 9, 9}),
		},
	}
	expected := [3][]int{
		{7, 0, 8},
		{0},
		{8, 9, 9, 9, 0, 0, 0, 1},
	}

	for i := 0; i < len(testcase); i++ {
		var temp = AddTwoNumbers(&testcase[i][0], &testcase[i][1])

		for j, val := range temp.ListToArray() {
			if val != expected[i][j] {
				t.Errorf("Incorrect result : expected %d got %d at test index %d", expected[i], temp.ListToArray(), i)
				break
			}
		}
	}

}
