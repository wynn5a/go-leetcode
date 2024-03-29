package solution_lcr_170

import "testing"

func TestReversePairs(t *testing.T) {
	testcases := []struct {
		record   []int
		expected int
	}{
		{[]int{7, 5, 6, 4}, 5},
		{[]int{1, 3, 2, 3, 1}, 4},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{4, 3, 2, 1}, 6},
	}

	for _, tc := range testcases {
		got := reversePairs(tc.record)
		if got != tc.expected {
			t.Errorf("record: %v, expected: %v, got: %v", tc.record, tc.expected, got)
		}
	}
}
