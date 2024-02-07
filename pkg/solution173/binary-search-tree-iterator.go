package solution173

import "github.com/wynn5a/go-leetcode/pkg/util"

type BSTIterator struct {
	Stack []*util.TreeNode
}

func Constructor(root *util.TreeNode) *BSTIterator {
	iterator := &BSTIterator{}
	iterator.leftmostInorder(root)
	return iterator
}

// leftmostInorder iterates over all the elements in the BST starting from the smallest one.
func (this *BSTIterator) leftmostInorder(root *util.TreeNode) {
	for root != nil {
		this.Stack = append(this.Stack, root)
		root = root.Left
	}
}

func (this *BSTIterator) Next() int {
	topmostNode := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]

	if topmostNode.Right != nil {
		this.leftmostInorder(topmostNode.Right)
	}

	return topmostNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}
