package string

///LeetCode Problems : Longest Common Prefix
///Link : https://leetcode.com/problems/longest-common-prefix/

func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	var result string = ""
	index := 0

	for {
		for _, str := range strs {
			if index >= len(str) {
				return result
			}
			if str[index] != strs[0][index] {
				return result
			}
		}
		result += string(strs[0][index])
		index++
	}
}
