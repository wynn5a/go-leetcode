package solution189

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, expected: []int{3, 99, -1, -100}},
	}

	for _, tc := range testCases {
		testNums := make([]int, len(tc.nums))
		copy(testNums, tc.nums)
		rotate(testNums, tc.k)
		if !reflect.DeepEqual(testNums, tc.expected) {
			t.Errorf("rotate(%v, %v): expected %v, but got %v",
				tc.nums, tc.k, tc.expected, testNums)
		}
	}
}
