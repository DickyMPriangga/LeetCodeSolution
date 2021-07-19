package main

import (
	"fmt"

	tp "github.com/DickyMPriangga/LeetCodeSolution/array/removeDuplicates"
	//lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"
)

func main() {
	test1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	test2 := tp.RemoveDuplicates(test1)
	fmt.Println(test1)
	fmt.Println(test2)
}
