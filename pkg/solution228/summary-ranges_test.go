package solution228

import (
	"reflect"
	"strconv"
	"testing"
)

func TestCases(t *testing.T) {
	cases := []struct {
		input  []int
		output []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{}, []string{}},
		{[]int{1}, []string{"1"}},
	}

	for i, c := range cases {
		t.Run("summaryRanges"+strconv.Itoa(i), func(t *testing.T) {
			result := summaryRanges(c.input)
			if !reflect.DeepEqual(result, c.output) {
				t.Errorf("summaryRanges(%v) = %v; expect %v", c.input, result, c.output)
			}
		})
	}
}
