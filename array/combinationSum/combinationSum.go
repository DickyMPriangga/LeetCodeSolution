package array

import "sort"

var endRes [][]int

func CombinationSum(candidates []int, target int) [][]int {
	currentRes := make([]int, 0)
	endRes = make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	CombinationRecur(candidates, target, currentRes, 0)
	return endRes
}

func CombinationRecur(candidates []int, target int, currentRes []int, current int) {

	if target == 0 {
		temp := make([]int, len(currentRes))
		copy(temp, currentRes)
		endRes = append(endRes, temp)
		return
	}

	for current < len(candidates) && target-candidates[current] >= 0 {
		currentRes = append(currentRes, candidates[current])
		CombinationRecur(candidates, target-candidates[current], currentRes, current)
		current++
		currentRes = currentRes[:len(currentRes)-1]
	}
}
