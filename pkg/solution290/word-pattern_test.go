package solution290

import "testing"

func TestCases(t *testing.T) {
	cases := []struct {
		pattern string
		s       string
		expect  bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"", "", true},
		{"a", "dog", true},
		{"a", "dog cat", false},
	}

	for _, c := range cases {
		result := wordPattern(c.pattern, c.s)
		if result != c.expect {
			t.Errorf("wordPattern(%v, %v) = %v; expect %v", c.pattern, c.s, result, c.expect)
		}
	}
}
