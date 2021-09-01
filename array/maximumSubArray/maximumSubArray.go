package array

func MaximumSubArray(nums []int) int {
	result := -100000
	current := 0

	for i := 0; i < len(nums); i++ {
		current += nums[i]
		if result < current {
			result = current
		}
		if current < 0 {
			current = 0
		}
	}

	return result
}
