package string

func FindSubstring(s string, words []string) []int {
	result := []int{}
	substr := make(map[string]int)
	lenWord := 0
	lenWordCnt := 0

	for _, val := range words {
		substr[val]++
		lenWordCnt++
		lenWord = len(val)
	}

	for i := 0; i < len(s)-(lenWordCnt*lenWord)+1; i++ {
		if _, ok := substr[s[i:i+lenWord]]; ok {
			testStr := s[i : i+(lenWordCnt*lenWord)]
			tempSubStr := make(map[string]int)
			for k, v := range substr {
				tempSubStr[k] = v
			}
			tempLenWordCnt := lenWordCnt
			for j := 0; j < (lenWordCnt * lenWord); j += lenWord {
				subSubStr := testStr[j : j+lenWord]
				if val, ok := tempSubStr[subSubStr]; ok && val > 0 {
					tempSubStr[subSubStr]--
					tempLenWordCnt--
				}
			}

			if tempLenWordCnt == 0 {
				result = append(result, i)
			}
		}
	}

	return result
}
