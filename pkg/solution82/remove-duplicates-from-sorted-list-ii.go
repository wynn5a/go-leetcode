package solution82

import "github.com/wynn5a/go-leetcode/pkg/util"

// deleteDuplicates removes all duplicates from the list and returns the head of the list after removal of duplicates in sorted order.
// It does not modify the original list.
func deleteDuplicates(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// create a dummy node to handle the case where the first node is a duplicate
	dummy := &util.ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			// skip all duplicates
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}
		curr = curr.Next
	}

	return dummy.Next
}
