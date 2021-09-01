package array

///LeetCode Problems : Container With Most Water
///Link : https://leetcode.com/problems/container-with-most-water/

func MaxArea(height []int) int {
	var result int = 0
	var waterLevel int = 1

	i := 0
	j := len(height) - 1

	for {
		for {
			if i == j || j < i {
				return result
			} else if height[i] < waterLevel {
				i++
			} else if height[j] < waterLevel {
				j--
			} else {
				break
			}
		}

		tempResult := (j - i) * waterLevel

		if tempResult > result {
			result = tempResult
		}

		waterLevel++
	}
}
