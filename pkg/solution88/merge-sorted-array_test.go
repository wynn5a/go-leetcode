package solution88

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	// Test cases
	testCases := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			expect: []int{1},
		},
		{
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			expect: []int{1},
		},
	}

	for _, tc := range testCases {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		if !reflect.DeepEqual(tc.nums1, tc.expect) {
			t.Errorf("Expected: %v, but got: %v", tc.expect, tc.nums1)
		}
	}
}
