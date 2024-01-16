package solution92

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		head     *util.ListNode
		left     int
		right    int
		expected *util.ListNode
	}{
		{
			name:     "Test 1: Reverse middle",
			head:     &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 4, Next: &util.ListNode{Val: 5}}}}},
			left:     2,
			right:    4,
			expected: &util.ListNode{Val: 1, Next: &util.ListNode{Val: 4, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 5}}}}},
		},
		{
			name:     "Test 2: Reverse entire",
			head:     &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3}}},
			left:     1,
			right:    3,
			expected: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 1}}},
		},
		{
			name:     "Test 3: Reverse one element",
			head:     &util.ListNode{Val: 5},
			left:     1,
			right:    1,
			expected: &util.ListNode{Val: 5},
		},
		{
			name:     "Test 4: Reverse first two",
			head:     &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3}}},
			left:     1,
			right:    2,
			expected: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 1, Next: &util.ListNode{Val: 3}}},
		},
		{
			name:     "Test 5: Reverse last two",
			head:     &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3}}},
			left:     2,
			right:    3,
			expected: &util.ListNode{Val: 1, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 2}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reversed := reverseBetween(tt.head, tt.left, tt.right)
			if !compareLinkedList(reversed, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, reversed)
			}
		})
	}
}

func compareLinkedList(l1, l2 *util.ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
