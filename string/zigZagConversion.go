package string

///LeetCode Problems : ZigZag Conversion
///Link : https://leetcode.com/problems/zigzag-conversion/

func Convert(s string, numRows int) string {
	var result string

	if numRows == 1 {
		return s
	}

	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < len(s); j += (2 * (numRows - 1)) {
				result += string(s[j])
			}
		} else {
			inverse := true
			for j := i; j < len(s); {
				if inverse {
					result += string(s[j])
					j += 2 * (numRows - i - 1)
				} else {
					result += string(s[j])
					j += 2 * i
				}
				inverse = !inverse
			}
		}
	}

	return result
}
