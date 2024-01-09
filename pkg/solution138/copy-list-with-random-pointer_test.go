package solution138

import "testing"

func TestCase1(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n1.Next = n2
	n1.Random = n2
	n2.Random = n2
	expected := n1
	actual := copyRandomList(n1)
	if !sameList(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func sameList(n1, n2 *Node) bool {
	if n1 == nil {
		return n2 == nil
	}
	if n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}
	if n1.Random != nil && n2.Random != nil {
		if n1.Random.Val != n2.Random.Val {
			return false
		}
	}
	return sameList(n1.Next, n2.Next)
}
