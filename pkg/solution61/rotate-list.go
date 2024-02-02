package solution61

import "github.com/wynn5a/go-leetcode/pkg/util"

func rotateRight(head *util.ListNode, k int) *util.ListNode {
	if head == nil || k == 0 {
		return head
	}

	// 1. Calculate the length of the list
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// 2. Calculate the actual number of rotations needed
	k = k % length

	// 3. Find the new head
	newHead := head
	for i := 0; i < length-k-1; i++ {
		newHead = newHead.Next
	}

	// 4. Reconnect the list
	tail.Next = head
	head = newHead.Next
	newHead.Next = nil

	return head
}
