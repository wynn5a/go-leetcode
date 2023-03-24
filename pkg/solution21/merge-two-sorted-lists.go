package solution21

import "github.com/wynn5a/go-leetcode/pkg/util"

func mergeTwoLists(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list2 != nil && list1 == nil {
		return list2
	}

	var rst *util.ListNode
	var tmp *util.ListNode

	for list1 != nil || list2 != nil {
		value := 0

		if list1 == nil {
			value = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			value = list1.Val
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			value = list1.Val
			list1 = list1.Next
		} else {
			value = list2.Val
			list2 = list2.Next
		}

		if rst == nil {
			rst = util.NewListNode(value)
			tmp = rst
		} else {
			tmp.Next = util.NewListNode(value)
			tmp = tmp.Next
		}
	}
	return rst
}
