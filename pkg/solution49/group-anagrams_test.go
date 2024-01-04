package solution49

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestCase(t *testing.T) {
	cases := []struct {
		input    []string
		expected [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			input: []string{""},
			expected: [][]string{
				{""},
			},
		},
		{
			input: []string{"a"},
			expected: [][]string{
				{"a"},
			},
		},
	}

	for _, c := range cases {
		actual := groupAnagrams(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d groups, got %d", len(c.expected), len(actual))
		}

		for _, group := range actual {
			checked := false
			m := make(map[int]bool)
			for i, strs := range c.expected {
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
				t.Errorf("Expected group %v to be one of %v", group, c.expected)
			}
		}
	}
}
