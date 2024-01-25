package solution148

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"strconv"
	"testing"
)

func TestSortList(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{-1, 5, 3, 2, 1, 0}, []int{-1, 0, 1, 2, 3, 5}},
	}

	for i, test := range cases {
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			result := sortList(util.NewListNodes(test.input))
			if !equal(result, util.NewListNodes(test.output)) {
				t.Errorf("sortList(%v) = %v; want %v", test.input, util.ListNodeToArray(result), test.output)
			}
		})
	}
}

func equal(result *util.ListNode, r *util.ListNode) bool {
	for result != nil && r != nil {
		if result.Val != r.Val {
			return false
		}
		result = result.Next
		r = r.Next
	}
	return result == nil && r == nil
}
