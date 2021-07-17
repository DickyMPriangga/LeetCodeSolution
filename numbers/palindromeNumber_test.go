package numbers

import "testing"

func TestIsPalindrome(t *testing.T) {
	testcase := []int{
		121,
		-121,
		10,
		-101,
	}

	expected := []bool{
		true,
		false,
		false,
		false,
	}

	for i := 0; i < len(testcase); i++ {
		result := IsPalindrome(testcase[i])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %t got %t at test %d index %d", expected[i], result, testcase[i], i)
		}
	}
}
