package solution19

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
)

func removeNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	if head == nil {
		return nil
	}

	dummy := &util.ListNode{Next: head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next

}
