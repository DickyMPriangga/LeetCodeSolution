package linkedList

import (
	"testing"

	lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"
)

//Testcase not finished yet
func TestMergeTwoList(t *testing.T) {
	testcase := [][2][]int{
		{{1, 2, 4}, {1, 3, 4}},
		{{}, {}},
		{{}, {0}},
	}

	expected := [][]int{
		{1, 1, 2, 3, 4, 4},
		{},
		{0},
	}

	for i := 0; i < len(testcase); i++ {
		list1 := lList.ListFromArray(testcase[i][0])
		list2 := lList.ListFromArray(testcase[i][1])

		result := MergeTwoLists(&list1, &list2)

		for j, val := range result.ListToArray() {
			if val != expected[i][j] {
				t.Errorf("Incorrect result : expected %d got %d at test index %d", expected[i], result.ListToArray(), i)
				break
			}
		}
	}
}
