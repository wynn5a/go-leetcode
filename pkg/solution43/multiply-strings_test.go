package solution43

import "testing"

func TestMultiply(t *testing.T) {
	testCases := []struct {
		num1     string
		num2     string
		expected string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"0", "123", "0"},
		{"999", "999", "998001"},
	}

	for _, tc := range testCases {
		result := multiply(tc.num1, tc.num2)
		if result != tc.expected {
			t.Errorf("multiply(%q, %q) expected %q but got %q", tc.num1, tc.num2, tc.expected, result)
		}
	}
}
