package string

import "strconv"

func LetterCombinations(digits string) []string {
	digitMap := make(map[int]string)

	digitMap[0] = ""
	digitMap[1] = ""
	digitMap[2] = "abc"
	digitMap[3] = "def"
	digitMap[4] = "ghi"
	digitMap[5] = "jkl"
	digitMap[6] = "mno"
	digitMap[7] = "pqrs"
	digitMap[8] = "tuv"
	digitMap[9] = "wxyz"

	var result []string
	temp := make([]int, len(digits))

	for {
		end := true
		for i := 0; i < len(temp); i++ {
			currDigits, _ := strconv.Atoi(string(digits[i]))
			if temp[i] < len(digitMap[currDigits]) {
				end = false
				break
			} else if i != len(temp)-1 {
				temp[i] = temp[i] % len(digitMap[currDigits])
				temp[i+1] += 1
			}
		}

		if end {
			return result
		}

		var tempStr string
		for i := 0; i < len(temp); i++ {
			currDigits, _ := strconv.Atoi(string(digits[i]))
			tempStr += string(digitMap[currDigits][temp[i]])
		}

		result = append(result, tempStr)
		temp[0]++
	}
}
