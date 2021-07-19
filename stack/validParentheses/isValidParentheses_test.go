package stack

import "testing"

func TestIsValid(t *testing.T) {
	testcase := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	expected := []bool{
		true,
		true,
		false,
		false,
		true,
	}

	for i := 0; i < len(testcase); i++ {
		result := IsValid(testcase[i])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %t got %t at test %s index %d", expected[i], result, testcase[i], i)
		}
	}
}
