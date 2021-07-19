package array

///LeetCode Problems : Remove Duplicates from Sorted Array
///Link : https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i := 0
	j := i + 1

	for j < len(nums) {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}

	return i + 1
}
