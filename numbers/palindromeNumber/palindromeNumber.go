package numbers

import "fmt"

///LeetCode Problems : Palindrome Number
///Link : https://leetcode.com/problems/palindrome-number/

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strInt := fmt.Sprint(x)

	var result string
	for _, v := range strInt {
		result = string(v) + result
	}

	return result == strInt
}
