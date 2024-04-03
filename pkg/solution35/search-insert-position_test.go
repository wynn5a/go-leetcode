package solution35

import "testing"

func TestCases(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
	}

	for _, c := range cases {
		if res := searchInsert(c.nums, c.target); res != c.expected {
			t.Errorf("Expected %d, but got %d", c.expected, res)
		}
	}
}
