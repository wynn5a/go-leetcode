package solution239

import (
	"reflect"
	"testing"
)

func TestCases(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{
			nums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:      3,
			expect: []int{3, 3, 5, 5, 6, 7},
		},
		{
			nums:   []int{1, 3, 1, 2, 0, 5},
			k:      3,
			expect: []int{3, 3, 2, 5},
		},
		{
			nums:   []int{1, 3, 1, 2, 0, 5},
			k:      6,
			expect: []int{5},
		},
	}

	for _, c := range cases {
		result := maxSlidingWindow(c.nums, c.k)
		if !reflect.DeepEqual(result, c.expect) {
			t.Errorf("maxSlidingWindow(%v, %d) = %v; expect %v", c.nums, c.k, result, c.expect)
		}
	}
}
