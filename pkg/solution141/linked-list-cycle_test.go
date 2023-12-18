package solution141

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestNoCycle(t *testing.T) {
	// 测试用例1：空链表
	assert(t, hasCycle(nil) == false, "there should no cycle in an empty list")

	// 测试用例2：没有环的链表
	list1 := &util.ListNode{Val: 3, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Next: &util.ListNode{Val: -1}}}}
	assert(t, hasCycle(list1) == false, "should no cycle in %v", util.ListNodeToArray(list1))
}

func TestHasCycle_HeadCycle(t *testing.T) {
	head := &util.ListNode{Val: 1}
	head.Next = head
	result := hasCycle(head)
	if !result {
		t.Errorf("Expected 'true' for a list with cycle at head, got 'false'")
	}
}

func TestHasCycle_MidCycle(t *testing.T) {
	head := &util.ListNode{Val: 1}
	head.Next = &util.ListNode{Val: 2}
	mid := head.Next
	head.Next.Next = &util.ListNode{Val: 3}
	head.Next.Next.Next = mid
	result := hasCycle(head)
	if !result {
		t.Errorf("Expected 'true' for a list with cycle in the middle, got 'false'")
	}
}

func assert(t *testing.T, condition bool, message string, args ...interface{}) {
	if !condition {
		t.Errorf(message, args...)
	}
}
