package solution82

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	cases := []struct {
		input  *util.ListNode
		output *util.ListNode
	}{
		{
			input:  util.NewListNodes([]int{1, 2, 3, 3, 4, 4, 5}),
			output: util.NewListNodes([]int{1, 2, 5}),
		},
		{
			input:  util.NewListNodes([]int{1, 1, 1, 2, 3}),
			output: util.NewListNodes([]int{2, 3}),
		},
		{
			input:  util.NewListNodes([]int{1, 1, 1}),
			output: nil,
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			output := deleteDuplicates(c.input)
			if !sameListNodes(output, c.output) {
				t.Errorf("Expected %v, got %v", c.output, output)
			}
		})
	}
}

func sameListNodes(l1 *util.ListNode, l2 *util.ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}
