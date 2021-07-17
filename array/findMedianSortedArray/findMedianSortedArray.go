package array

import "sort"

///LeetCode Problems : Median of Two Sorted Arrays
///Link : https://leetcode.com/problems/median-of-two-sorted-arrays/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums3 []int

	n3 := nums3[:]
	n3 = append(n3, nums1...)
	n3 = append(n3, nums2...)

	sort.Ints(n3)

	if len(n3)%2 == 0 {
		median1 := float64(n3[len(n3)/2])
		median2 := float64(n3[len(n3)/2-1])
		return (median1 + median2) / 2
	} else {
		return float64(n3[int(len(n3)/2)])
	}
}
