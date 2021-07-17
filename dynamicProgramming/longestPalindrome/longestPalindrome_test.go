package dynamicProgramming

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCase := []string{"babad", "cbbd", "a", "ac"}
	expected := []string{"bab", "bb", "a", "a"}

	for i := 0; i < len(testCase); i++ {
		test := LongestPalindrome(testCase[i])
		if test != expected[i] {
			t.Errorf("Incorrect result : expected %s got %s", expected[i], test)
		}
	}
}

func TestBruteLongestPalindrome(t *testing.T) {
	testCase := []string{"babad", "cbbd", "a", "ac"}
	expected := []string{"bab", "bb", "a", "a"}

	for i := 0; i < len(testCase); i++ {
		test := BruteLongestPalindrome(testCase[i])
		if test != expected[i] {
			t.Errorf("Incorrect result : expected %s got %s", expected[i], test)
		}
	}
}
