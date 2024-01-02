package solution242

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		str1, str2 string
		expected   bool
	}{
		{"listen", "silent", true},
		{"hello", "world", false},
		{"abb", "bab", true},
		{"abc", "cba", true},
		{"aabbcc", "ccbbaa", true},
		{"abc", "abcd", false},
	}

	for _, test := range tests {
		result := isAnagram(test.str1, test.str2)
		if result != test.expected {
			t.Errorf("For str1 = %s, str2 = %s, expected = %t, got = %t", test.str1, test.str2, test.expected, result)
		}
	}
}
