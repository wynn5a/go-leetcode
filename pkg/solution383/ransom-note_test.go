package solution383

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"", "abc", true},
		{"abc", "", false},
		{"abc", "cba", true},
		{"aaa", "aa", false},
	}

	for _, tt := range tests {
		actual := canConstruct(tt.ransomNote, tt.magazine)
		if actual != tt.expected {
			t.Errorf("ransomNote: %s, magazine: %s, expected: %t, actual: %t",
				tt.ransomNote, tt.magazine, tt.expected, actual)
		}
	}
}
