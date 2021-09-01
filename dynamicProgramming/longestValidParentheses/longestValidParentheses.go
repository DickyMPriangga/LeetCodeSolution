package dynamicProgramming

func LongestValidParentheses(s string) int {
	var symbolStack []int
	var symbolMap map[int]int = make(map[int]int)
	result := 0

	for i, val := range s {
		if len(symbolStack) == 0 {
			if val == '(' {
				symbolStack = append(symbolStack, i)
				symbolMap[i] = -1
				continue
			} else {
				continue
			}
		}

		if string(val) == "(" {
			symbolMap[i] = -1
			symbolStack = append(symbolStack, i)
		} else if len(symbolStack) > 0 {
			symbolMap[symbolStack[len(symbolStack)-1]] = i - symbolStack[len(symbolStack)-1] + 1
			temp := i - symbolStack[len(symbolStack)-1] + 1
			if temp > result {
				result = temp
			}
			symbolStack = symbolStack[:len(symbolStack)-1]
		}
	}

	curr := 0
	for key, val := range symbolMap {
		if val == 2 {
			if curr == -1 {
				curr = key
			} else {
				symbolMap[curr] += 2
				if symbolMap[curr] > result {
					result = symbolMap[curr]
				}
				delete(symbolMap, key)
			}
		} else {
			curr = -1
		}
	}

	return result
}

func LongestValidParentheses2(s string) int {
	var symbolStack []string
	result := 0
	//temp := 0

	for _, val := range s {
		if len(symbolStack) == 0 {
			if string(val) == "(" {
				symbolStack = append(symbolStack, string(val))
			}
			continue
		}

		if string(val) == "(" {
			symbolStack = append(symbolStack, string(val))
		} else {
			symbolStack = symbolStack[:len(symbolStack)-1]
			result += 2
		}
	}

	return result
}
