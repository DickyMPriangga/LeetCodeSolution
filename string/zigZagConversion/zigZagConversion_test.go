package string

import "testing"

func TestConvert(t *testing.T) {
	testcase1 := []string{
		"PAYPALISHIRING",
		"PAYPALISHIRING",
		"A",
	}
	testcase2 := []int{
		3,
		4,
		1,
	}

	expected := []string{
		"PAHNAPLSIIGYIR",
		"PINALSIGYAHRPI",
		"A",
	}

	for i := 0; i < len(testcase1); i++ {
		result := Convert(testcase1[i], testcase2[i])
		if result != expected[i] {
			t.Errorf("Incorrect result : expected %s got %s at test %s,%d index %d", expected[i], result, testcase1[i], testcase2[i], i)
		}
	}
}
