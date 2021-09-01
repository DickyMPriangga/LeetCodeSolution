package array

func TrapRainWater(height []int) int {
	result := 0

	for i := 0; i < len(height); i++ {
		if i == 0 && height[i] == 0 {
			continue
		}

		nextHigh := -1
		areaHeight := 0

		for j := i + 1; j < len(height); j++ {
			if height[j] >= height[i] {
				nextHigh = j
				break
			} else {
				areaHeight += height[j]
			}
		}

		if nextHigh != -1 {
			result += ((nextHigh - i - 1) * height[i]) - areaHeight
			i = nextHigh - 1
		}
	}

	return result
}
