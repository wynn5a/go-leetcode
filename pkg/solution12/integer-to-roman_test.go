package solution12

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{42, "XLII"},
		{3999, "MMMCMXCIX"},
		{0, ""},
		{-1, ""},
		{4000, ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input %d should convert to %s", tc.input, tc.expected), func(t *testing.T) {
			result := intToRoman(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s but got %s", tc.expected, result)
			}
		})
	}
}
