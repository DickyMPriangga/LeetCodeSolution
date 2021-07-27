package array

func SearchRotatedSortArray(nums []int, target int) int {

	return binarySearch(nums, 0, len(nums)-1, target)

}

func binarySearch(nums []int, left, right, target int) int {
	if right < left {
		return -1
	}

	if right-left < 2 {
		if nums[right] == target {
			return right
		} else if nums[left] == target {
			return left
		} else {
			return -1
		}
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
		if target <= nums[right] {
			return binarySearch(nums, mid+1, right, target)
		} else {
			return binarySearch(nums, left, mid-1, target)
		}
	} else if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
		if target >= nums[left] {
			return binarySearch(nums, left, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, right, target)
		}
	} else {
		if nums[right] < nums[mid] {
			if target < nums[mid] && target >= nums[left] {
				return binarySearch(nums, left, mid-1, target)
			} else {
				return binarySearch(nums, mid+1, right, target)
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				return binarySearch(nums, mid+1, right, target)
			} else {
				return binarySearch(nums, left, mid-1, target)
			}
		}
	}
}
