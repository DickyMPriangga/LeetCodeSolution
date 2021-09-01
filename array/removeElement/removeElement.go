package array

func RemoveElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}

	i := 0
	j := i + 1

	for j < len(nums) {
		if nums[i] == val && nums[j] != val {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
		} else if nums[i] == val && nums[j] == val {
			j++
		} else {
			i++
			j++
		}
	}

	if nums[i] != val {
		i++
	}

	return i
}
