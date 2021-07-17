package string

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testcase := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
	}

	expected := []int{
		3, 1, 3, 0,
	}

	for i := 0; i < len(testcase); i++ {
		result := LengthOfLongestSubstring(testcase[i])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %d got %d at test %s index %d", expected[i], result, testcase[i], i)
		}
	}
}
