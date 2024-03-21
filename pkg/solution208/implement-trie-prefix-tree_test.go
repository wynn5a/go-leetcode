package solution208

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("expected true, but got false")
	}
	if trie.Search("app") {
		t.Errorf("expected false, but got true")
	}
	if !trie.StartsWith("app") {
		t.Errorf("expected true, but got false")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("expected true, but got false")
	}
}

func TestTrieStartsWith(t *testing.T) {
	trie := Constructor()
	trie.Insert("hello")
	trie.Insert("world")

	tests := []struct {
		prefix string
		want   bool
	}{
		{"he", true},
		{"hello", true},
		{"helloworld", false},
		{"wo", true},
		{"world", true},
		{"wonderful", false},
	}

	for _, test := range tests {
		got := trie.StartsWith(test.prefix)
		if got != test.want {
			t.Errorf("For prefix %q: Expected %t, but got %t", test.prefix, test.want, got)
		}
	}
}
