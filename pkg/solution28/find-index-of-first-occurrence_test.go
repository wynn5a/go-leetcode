package solution28

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"mississippi", "issip", 4},
		{"leetcode", "leetcode1", -1},
	}

	for _, testCase := range testCases {
		result := strStr(testCase.haystack, testCase.needle)
		if result != testCase.expected {
			t.Fatalf("strStr('%s', '%s') returns %d, expected %d", testCase.haystack, testCase.needle, result, testCase.expected)
		}
	}
}
