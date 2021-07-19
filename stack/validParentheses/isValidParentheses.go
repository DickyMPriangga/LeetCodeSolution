package stack

///LeetCode Problems : Valid Parentheses
///Link : https://leetcode.com/problems/valid-parentheses/

func IsValid(s string) bool {
	var symbolStack []string

	mapSymbol := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	for _, val := range s {
		if len(symbolStack) == 0 {
			symbolStack = append(symbolStack, string(val))
			continue
		}

		if string(val) != mapSymbol[symbolStack[len(symbolStack)-1]] {
			symbolStack = append(symbolStack, string(val))
		} else {
			symbolStack = symbolStack[:len(symbolStack)-1]
		}
	}

	if len(symbolStack) > 0 {
		return false
	} else {
		return true
	}
}
