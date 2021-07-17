package string

import "strings"

///LeetCode Problems : Longest Substring Without Repeating Characters
///Link : https://leetcode.com/problems/longest-substring-without-repeating-characters/

func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var resultString string
	var result = 0

	for i := 0; i < len(s); i++ {
		resultString = ""
		for j := i; j < len(s); j++ {
			if strings.Contains(resultString, string(s[j])) {
				break
			} else {
				resultString += string(s[j])
			}
		}

		if len(resultString) > result {
			result = len(resultString)
		}
	}

	return result
}
