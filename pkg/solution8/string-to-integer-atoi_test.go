package solution8

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"42", 42},
		{"-42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"   -42", -42},
		{"+1", 1},
		{"-2147483649", -2147483648},
	}

	for _, test := range tests {
		result := myAtoi(test.input)
		if result != test.expected {
			t.Errorf("myAtoi(%q) returned %d, expected %d", test.input, result, test.expected)
		}
	}
}
