package array

import (
	"sort"
)

///LeetCode Problems : 3Sum
///Link : https://leetcode.com/problems/3sum/

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}

		j := i + 1
		k := len(nums) - 1

		for {
			if j >= k {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for {
					j++
					if j >= k || nums[j] != nums[j-1] {
						break
					}
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				for {
					j++
					if j >= k || nums[j] != nums[j-1] {
						break
					}
				}
			} else {
				for {
					k--
					if j >= k || nums[k] != nums[k+1] {
						break
					}
				}
			}
		}
	}

	return result
}
