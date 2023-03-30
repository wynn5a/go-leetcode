package solution24

import "github.com/wynn5a/go-leetcode/pkg/util"

func swapPairs(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &util.ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next
		prev.Next, first.Next, second.Next = second, second.Next, first
		prev = first
	}
	return dummy.Next
}
