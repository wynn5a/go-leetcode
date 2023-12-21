package solution209

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		target int
		nums   []int
		result int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	for _, tc := range testCases {
		res := minSubArrayLen(tc.target, tc.nums)
		if res != tc.result {
			t.Errorf("Expected %d, but got %d", tc.result, res)
		}
	}
}
