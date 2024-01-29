package solution155

import "testing"

func TestCases(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	m := minStack.GetMin() // 返回 -3
	if m != -3 {
		t.Errorf("GetMin() = %d; want %d", m, -3)
	}
	minStack.Pop()
	top := minStack.Top() // 返回 0
	if top != 0 {
		t.Errorf("Top() = %d; want %d", top, 0)
	}
	m = minStack.GetMin() // 返回 -2
	if m != -2 {
		t.Errorf("GetMin() = %d; want %d", m, -2)
	}
	minStack.Push(1)
	top = minStack.Top() // 返回 1
	if top != 1 {
		t.Errorf("Top() = %d; want %d", top, 1)
	}
	m = minStack.GetMin() // 返回 -2
	if m != -2 {
		t.Errorf("GetMin() = %d; want %d", m, -2)
	}
}
