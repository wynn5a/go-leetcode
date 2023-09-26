package solution151

import (
	"strings"
	"testing"
)

func TestCases(t *testing.T) {
	args := []struct {
		s        string
		expected string
	}{
		{s: "the sky is blue", expected: "blue is sky the"},
		{s: "  hello world  ", expected: "world hello"},
		{s: "a good   example", expected: "example good a"},
	}

	for _, arg := range args {
		words := reverseWords(arg.s)
		if strings.Compare(arg.expected, words) != 0 {
			t.Errorf("Test failed for string: %s, expected:{%s}, actual:{%s}", arg.s, arg.expected, words)
		}
	}
}
