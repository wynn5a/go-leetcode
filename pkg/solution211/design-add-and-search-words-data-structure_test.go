package solution211

import "testing"

func TestWordDictionary(t *testing.T) {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	tests := []struct {
		word   string
		expect bool
	}{
		{"pad", false},
		{"bad", true},
		{".ad", true},
		{"b..", true},
	}

	for _, test := range tests {
		if got := wd.Search(test.word); got != test.expect {
			t.Errorf("for word %q, expected %t but got %t", test.word, test.expect, got)
		}
	}
}
