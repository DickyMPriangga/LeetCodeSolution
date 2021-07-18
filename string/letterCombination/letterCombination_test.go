package string

import (
	"testing"
)

func TestLetterCombination(t *testing.T) {
	testcase := []string{
		"23",
		"",
		"2",
	}

	expected := [][]string{
		{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		{},
		{"a", "b", "c"},
	}

	for i := 0; i < len(testcase); i++ {
		result := LetterCombinations(testcase[i])
		for _, val := range result {
			if !contains(expected[i], val) {
				t.Errorf("Incorrect result : expected %s got %s at test %s index %d", expected[i], result, testcase[i], i)
			}
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
