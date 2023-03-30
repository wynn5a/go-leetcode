package solution24

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	// 构造链表
	head1 := &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 4}}}}
	expected1 := &util.ListNode{Val: 2, Next: &util.ListNode{Val: 1, Next: &util.ListNode{Val: 4, Next: &util.ListNode{Val: 3}}}}
	result1 := swapPairs(head1)
	if !isEqual(result1, expected1) {
		t.Errorf("Error: expected %v, but got %v", expected1, result1)
	}

	head2 := &util.ListNode{Val: 1}
	expected2 := &util.ListNode{Val: 1}
	result2 := swapPairs(head2)
	if !isEqual(result2, expected2) {
		t.Errorf("Error: expected %v, but got %v", expected2, result2)
	}

	head3 := &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2}}
	expected3 := &util.ListNode{Val: 2, Next: &util.ListNode{Val: 1}}
	result3 := swapPairs(head3)
	if !isEqual(result3, expected3) {
		t.Errorf("Error: expected %v, but got %v", expected3, result3)
	}

	head4 := &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3}}}
	expected4 := &util.ListNode{Val: 2, Next: &util.ListNode{Val: 1, Next: &util.ListNode{Val: 3}}}
	result4 := swapPairs(head4)
	if !isEqual(result4, expected4) {
		t.Errorf("Error: expected %v, but got %v", expected4, result4)
	}
}

// 判断两个链表是否相等
func isEqual(l1 *util.ListNode, l2 *util.ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
