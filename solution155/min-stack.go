package solution155

import (
	"math"
)

// MinStack use two linked lists or use two slices to store the min value
type MinStack struct {
	//Head *util.ListNode
	//Min  *util.ListNode

	Stack []int
	Min   []int
}

func Constructor() *MinStack {
	return &MinStack{}
}

func (this *MinStack) Push(val int) {
	//if this.Head == nil {
	//	this.Head = &util.ListNode{Val: val}
	//	this.Min = &util.ListNode{Val: val}
	//} else {
	//	this.Head = &util.ListNode{Val: val, Next: this.Head}
	//	if val < this.Min.Val {
	//		this.Min = &util.ListNode{Val: val, Next: this.Min}
	//	} else {
	//		this.Min = &util.ListNode{Val: this.Min.Val, Next: this.Min}
	//	}
	//}
	this.Stack = append(this.Stack, val)
	if len(this.Min) == 0 || val <= this.Min[len(this.Min)-1] {
		this.Min = append(this.Min, val)
	}
}

func (this *MinStack) Pop() {
	//if this.Head == nil {
	//	return
	//}
	//this.Min = this.Min.Next
	//this.Head = this.Head.Next
	if len(this.Stack) == 0 {
		return
	}
	if this.Stack[len(this.Stack)-1] == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	//return this.Head.Val
	if len(this.Stack) == 0 {
		return math.MinInt
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	//return this.Min.Val
	if len(this.Min) == 0 {
		return -1
	}
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
