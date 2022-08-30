package solution543

import "fmt"

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// max sum of left child and right child for every node
func diameterOfNode(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	fmt.Printf("current: %d, max: %d \n", root.Val, *max)

	left := diameterOfNode(root.Left, max)
	fmt.Printf("left: %d \n", left)
	right := diameterOfNode(root.Right, max)
	fmt.Printf("right: %d \n", right)

	result := maxOf(left, right) + 1
	if *max < left+right {
		*max = left + right
	}
	return result
}

func diameterOfBinaryTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var max *int
	max = new(int)
	*max = 0
	left := diameterOfNode(node.Left, max)
	right := diameterOfNode(node.Right, max)

	return maxOf(*max, left+right)
}

func maxOf(a, b int) int {
	if a < b {
		return b
	}
	return a
}
