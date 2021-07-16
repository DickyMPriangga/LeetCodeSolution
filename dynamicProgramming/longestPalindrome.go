package dynamicProgramming

func BruteLongestPalindrome(s string) string {
	var longest int = 0
	var longestString string = ""

	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			var result string
			for _, v := range s[i:j] {
				result = string(v) + result
			}

			if result == s[i:j] && len(result) > longest {
				longest = len(result)
				longestString = result
			}
		}
	}

	return longestString
}

func LongestPalindrome(s string) string {
	var result string
	var temp string

	for i := 0; i < len(s); i++ {
		temp = checkPalindrome(s, i, i)
		if len(temp) > len(result) {
			result = temp
		}
		temp = checkPalindrome(s, i, i+1)
		if len(temp) > len(result) {
			result = temp
		}
	}

	return result
}

func checkPalindrome(s string, l int, r int) string {
	if l < r || s == "" {
		return ""
	}

	var result string

	for {
		if l < 0 || r >= len(s) || s[l] != s[r] {
			break
		}

		result = s[l:r]
		l--
		r++
	}

	return result
}
