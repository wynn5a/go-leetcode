package solution14

import "testing"

func TestCases(t *testing.T) {
	cases := []struct {
		input  []string
		output string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "racecar", "car"}, ""},
		{[]string{"", ""}, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "a"}, "a"},
		{[]string{"a", "b"}, ""},
	}

	for _, c := range cases {
		result := longestCommonPrefix(c.input)
		if result != c.output {
			t.Errorf("longestCommonPrefix(%v) = %s; expect %s", c.input, result, c.output)
		}
	}
}
