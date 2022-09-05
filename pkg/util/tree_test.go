package util

import (
	"fmt"
	"testing"
)

func TestArray2TreeNode(t *testing.T) {
	input := "[1,null,2,3,4,5,6,7]"
	root := Array2TreeNode(input)
	visit(root)
}

func visit(root *TreeNode) {
	if root == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(root.Val)
	visit(root.Left)
	visit(root.Right)
}
