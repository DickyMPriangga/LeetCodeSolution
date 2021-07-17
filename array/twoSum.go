package array

func TwoSum(nums []int, target int) []int {
	var result = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}

	return result
}
