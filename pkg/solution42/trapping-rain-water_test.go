package solution42

import "testing"

func TestTrap(t *testing.T) {
	testCases := []struct {
		height []int
		want   int
	}{
		{[]int{}, 0},
		{[]int{2}, 0},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{2, 1, 2, 1, 2, 1}, 2},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	for _, tc := range testCases {
		got := trap(tc.height)
		if got != tc.want {
			t.Errorf("trap(%v) = %v; want %v", tc.height, got, tc.want)
		}
	}
}
