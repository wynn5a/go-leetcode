package solution2

import (
	"fmt"
	u "github.com/wynn5a/go-leetcode/pkg/util"
)

func addTwoNumbers(l1 *u.ListNode, l2 *u.ListNode) *u.ListNode {
	var rst, t *u.ListNode
	tmp := 0
	for l2 != nil || l1 != nil {
		var v1, v2 = 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		r := v1 + v2 + tmp
		if r >= 10 {
			tmp = r / 10
			r = r % 10
		} else {
			tmp = 0
		}

		fmt.Println(v1, v2, r, tmp)

		if rst == nil {
			rst = &u.ListNode{
				Val:  r,
				Next: nil,
			}
			t = rst
		} else {
			next := &u.ListNode{
				Val:  r,
				Next: nil,
			}
			t.Next = next
			t = t.Next
		}

	}

	if tmp == 1 {
		t.Next = &u.ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return rst
}
