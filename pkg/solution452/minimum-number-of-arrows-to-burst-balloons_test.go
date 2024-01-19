package solution452

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			expected: 2,
		},
		{
			input:    [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			expected: 4,
		},
		{
			input:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			expected: 2,
		},
		{
			input:    [][]int{{1, 2}},
			expected: 1,
		},
		{
			input:    [][]int{{2, 3}, {2, 3}},
			expected: 1,
		},
	}

	for _, c := range cases {
		actual := findMinArrowShots(c.input)
		if actual != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, actual)
		}
	}
}
