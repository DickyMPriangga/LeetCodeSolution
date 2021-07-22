package string

import "testing"

func TestFindSubstring(t *testing.T) {
	testcase1 := []string{
		"barfoothefoobarman",
		"wordgoodgoodgoodbestword",
		"barfoofoobarthefoobarman",
		"wordgoodgoodgoodbestword",
	}

	testcase2 := [][]string{
		{"foo", "bar"},
		{"word", "good", "best", "word"},
		{"bar", "foo", "the"},
		{"word", "good", "best", "good"},
	}

	expected := [][]int{
		{0, 9},
		{},
		{6, 9, 12},
		{8},
	}

	for i := 0; i < len(testcase1); i++ {
		result := FindSubstring(testcase1[i], testcase2[i])
		for j := 0; j < len(result); j++ {
			if result[j] != expected[i][j] {
				t.Errorf("Incorrect result : expected %d got %d at test %s,%s index %d", expected[i], result, testcase1[i], testcase2[i], i)
			}
		}
	}
}
