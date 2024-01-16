package solution92

import "github.com/wynn5a/go-leetcode/pkg/util"

func reverseBetween(head *util.ListNode, left int, right int) *util.ListNode {
	if left == right {
		return head
	}

	dummy := &util.ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}
