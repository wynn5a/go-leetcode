package solution49

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestCase(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	actual := groupAnagrams(input)
	if len(actual) != len(expected) {
		t.Errorf("Expected %d groups, got %d", len(expected), len(actual))
	}

	for _, group := range actual {
		checked := false
		m := make(map[int]bool)
		for i, strs := range expected {
			if m[i] {
				continue
			}
			if util.SameStringSlice(group, strs) {
				m[i] = true
				checked = true
				break
			}
		}
		if !checked {
			t.Errorf("Expected group %v to be one of %v", group, expected)
		}
	}
}
