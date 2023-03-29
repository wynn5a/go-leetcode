package solution31

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, tc := range testCases {
		nextPermutation(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.expected) {
			t.Errorf("for nums=%v, expected=%v but got=%v", tc.nums, tc.expected, tc.nums)
		}
	}
}
