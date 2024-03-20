package solution167

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tc := range testcases {
		got := twoSum(tc.numbers, tc.target)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("twoSum(%v, %v) = %v, expected %v", tc.numbers, tc.target, got, tc.expected)
		}
	}
}
