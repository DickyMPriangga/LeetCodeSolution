package numbers

import "strconv"

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	say := CountAndSay(n - 1)
	current := '0'
	currentCnt := 0
	result := ""

	for i, val := range say {
		if val != current {
			if current == '0' {
				current = val
			} else {
				result += strconv.Itoa(currentCnt)
				result += string(current)
				current = val
				currentCnt = 0
			}
		}

		if val == current {
			currentCnt++
		}

		if i == len(say)-1 {
			result += strconv.Itoa(currentCnt)
			result += string(current)
			current = val
			currentCnt = 0
		}
	}

	return result
}
