package string

import "testing"

func TestStrStr(t *testing.T) {
	testcase1 := []string{
		"hello",
		"aaaaa",
		"",
		"a",
	}

	testcase2 := []string{
		"ll",
		"bba",
		"",
		"a",
	}

	expected := []int{
		2,
		-1,
		0,
		0,
	}

	for i := 0; i < len(testcase1); i++ {
		result := StrStr(testcase1[i], testcase2[i])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %d got %d at test %s,%s index %d", expected[i], result, testcase1[i], testcase2[i], i)
		}
	}
}
