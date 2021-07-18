package array

import (
	"sort"
)

///LeetCode Problems : 3Sum Closest
///Link : https://leetcode.com/problems/3sum-closest/

func ThreeSumClosest(nums []int, target int) int {
	var result int = -1000
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1

		for {
			if j >= k {
				break
			}

			if nums[i]+nums[j]+nums[k] <= target {
				if AbsInt(nums[i]+nums[j]+nums[k]-target) < AbsInt(result-target) {
					result = nums[i] + nums[j] + nums[k]
				}
				j++
			} else {
				if AbsInt(nums[i]+nums[j]+nums[k]-target) < AbsInt(result-target) {
					result = nums[i] + nums[j] + nums[k]
				}
				k--
			}
		}
	}

	return result
}

// Abs returns the absolute value of x.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
