package main

import (
	"fmt"

	api "github.com/DickyMPriangga/LeetCodeSolution/array/seachRotatedSortArray"
	//lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"
)

func main() {
	nums := []int{1, 3, 5}
	target := 1

	fmt.Println(api.SearchRotatedSortArray(nums, target))
}
