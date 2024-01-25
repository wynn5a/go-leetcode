package solution148

import (
	"fmt"
	"github.com/wynn5a/go-leetcode/pkg/util"
)

func sortList(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find the middle of the list
	middle := findMiddle(head)

	// split the list into two halves
	right := sortList(middle.Next)
	middle.Next = nil
	left := sortList(head)

	// merge the two sorted lists
	node := merge(left, right)
	fmt.Printf("merged: %v\n", util.ListNodeToArray(node))
	return node
}

func merge(left *util.ListNode, right *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dummy

	fmt.Printf("left: %v\n", util.ListNodeToArray(left))
	fmt.Printf("right: %v\n", util.ListNodeToArray(right))

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}

// findMiddle finds the middle of the list and returns the middle node
func findMiddle(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
