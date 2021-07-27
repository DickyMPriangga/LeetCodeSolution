package array

func SearchRangeSortedArray(nums []int, target int) []int {
	targetPos := binarySearch(nums, 0, len(nums)-1, target)

	if target == -1 {
		return []int{-1, -1}
	} else {
		p1, p2 := targetPos, targetPos
		for {
			f1, f2 := true, true
			if p1 > 0 && nums[p1-1] == target {
				p1--
			} else {
				f1 = false
			}
			if p2 < len(nums)-1 && nums[p2+1] == target {
				p2++
			} else {
				f2 = false
			}

			if !f1 && !f2 {
				break
			}
		}

		return []int{p1, p2}
	}
}

func binarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	} else if target > nums[mid] {
		return binarySearch(nums, mid+1, right, target)
	} else {
		return binarySearch(nums, left, mid-1, target)
	}
}
