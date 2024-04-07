package solution1071

import "testing"

func TestCases(t *testing.T) {
	cases := []struct {
		str1 string
		str2 string
		want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
		{"ABCDEF", "ABC", ""},
	}

	for _, c := range cases {
		if got := gcdOfStrings(c.str1, c.str2); got != c.want {
			t.Errorf("str1: %s, str2: %s, expected: %s, but got: %s", c.str1, c.str2, c.want, got)
		}
	}
}
