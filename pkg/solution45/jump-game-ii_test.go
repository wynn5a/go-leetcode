package solution45

import "testing"

func TestJump(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 1, 1, 1, 1}, 4},
		{[]int{0}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
	}

	for _, tc := range testCases {
		actual := jump(tc.nums)
		if actual != tc.expected {
			t.Errorf("jump(%v) = %d, expected %d", tc.nums, actual, tc.expected)
		}
	}
}
