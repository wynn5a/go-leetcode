package solution169

import "testing"
import "reflect"

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tc := range testCases {
		actual := majorityElement(tc.nums)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("findMajorityElement(%v): expected %v, got %v",
				tc.nums, tc.expected, actual)
		}
	}
}
