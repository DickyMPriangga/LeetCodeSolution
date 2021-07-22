package string

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack)-(len(needle)-1); i++ {
		if haystack[i] == needle[0] {
			j := 0
			k := len(needle) - 1
			str := haystack[i : i+len(needle)]
			flag := true
			for j <= k {
				if str[j] == needle[j] && str[k] == needle[k] {
					j++
					k--
				} else {
					flag = false
					break
				}
			}

			if flag {

				return i
			}
		}
	}

	return -1
}
