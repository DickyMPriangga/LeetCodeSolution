package numbers

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	testcase := []struct {
		input    int
		expected string
	}{
		{
			input:    1,
			expected: "1",
		},
		{
			input:    4,
			expected: "1211",
		},
		{
			input:    6,
			expected: "312211",
		},
	}

	for i, tc := range testcase {
		assert.Equal(t, tc.expected, CountAndSay(tc.input), "Testcase "+strconv.Itoa(i))
	}
}
