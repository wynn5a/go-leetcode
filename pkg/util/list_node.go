package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) (r *ListNode) {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func NewListNodes(nums []int) (r *ListNode) {
	var t *ListNode
	for _, v := range nums {
		if r == nil {
			r = &ListNode{
				Val:  v,
				Next: nil,
			}
			t = r
		} else {
			tmp := &ListNode{
				Val:  v,
				Next: nil,
			}
			t.Next = tmp
			t = t.Next
		}
	}
	return r
}

func ListNodeToArray(l *ListNode) []int {
	var rst []int

	for l != nil {
		rst = append(rst, l.Val)
		l = l.Next
	}

	return rst
}
