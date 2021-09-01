package main

import (
	api "github.com/DickyMPriangga/LeetCodeSolution/array/combinationSum"
	//lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"
)

func main() {
	candidates := []int{2, 3, 5}
	target := 8

	api.CombinationSum(candidates, target)
}
