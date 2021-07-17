package string

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testcase := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{""},
	}

	expected := []string{
		"fl",
		"",
		"",
	}

	for i := 0; i < len(testcase); i++ {
		result := LongestCommonPrefix(testcase[i])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %s got %s at test %s index %d", expected[i], result, testcase[i], i)
		}
	}
}
