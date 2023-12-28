package solution205

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
		{"", "", true},
		{"aba", "baa", false},
	}

	for _, test := range tests {
		got := isIsomorphic(test.s, test.t)
		if got != test.want {
			t.Errorf("isIsomorphic(%q, %q): got %v, want %v", test.s, test.t, got, test.want)
		}
	}
}
