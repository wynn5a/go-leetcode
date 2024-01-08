package solution86

import "github.com/wynn5a/go-leetcode/pkg/util"

// partition splits a linked list into two lists, one with nodes with value < x, one with nodes with value >= x
// and then connects the two lists
func partition(head *util.ListNode, x int) *util.ListNode {
	// create two lists, one for nodes with value < x, one for nodes with value >= x
	l1 := &util.ListNode{}
	l2 := &util.ListNode{}
	// create two pointers to the two lists
	p1 := l1
	p2 := l2
	for head != nil {
		if head.Val < x {
			// append to l1
			p1.Next = head
			p1 = p1.Next
		} else {
			// append to l2
			p2.Next = head
			p2 = p2.Next
		}
		// move head to next node
		head = head.Next
	}
	// connect l1 and l2
	p2.Next = nil
	p1.Next = l2.Next
	return l1.Next
}
