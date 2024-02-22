package solution25

import "github.com/wynn5a/go-leetcode/pkg/util"

func reverseKGroup(head *util.ListNode, k int) *util.ListNode {
	// 一个虚拟头节点来简化操作
	dummy := &util.ListNode{Next: head}
	prev, end := dummy, dummy

	for end.Next != nil {
		// 先走k步找到当前组的末尾
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 如果end为空，说明剩余节点不足k个，不需要翻转
		if end == nil {
			break
		}
		// 记录一下要翻转部分的头和尾，以及下个部分的头
		start := prev.Next
		next := end.Next
		// 切断end和next，准备翻转
		end.Next = nil
		// 翻转从start到end的部分
		prev.Next = reverse(start)
		// 翻转完成后，start成了这部分的末尾，连接到next
		start.Next = next
		// 移动prev和end到下一组的起始位置
		prev = start
		end = prev
	}
	return dummy.Next
}

func reverse(head *util.ListNode) *util.ListNode {
	var prev *util.ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}
