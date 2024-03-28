package solution76

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for _, test := range tests {
		if got := minWindow(test.s, test.t); got != test.expect {
			t.Errorf("for s=%q, t=%q, expected %q but got %q", test.s, test.t, test.expect, got)
		}
	}
}
