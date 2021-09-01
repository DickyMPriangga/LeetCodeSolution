package array

func SearchInsert(nums []int, target int) int {
	found, result := binarySearch(nums, 0, len(nums)-1, target)

	if found {
		return result
	} else {
		if target >= nums[result] {
			return result + 1
		} else {
			return result
		}
	}
}

func binarySearch(nums []int, left, right, target int) (bool, int) {
	mid := left + (right-left)/2

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return true, mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false, mid
}
