package solution653

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	var stack []*TreeNode
	var m = make(map[int]bool)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := m[k-root.Val]; ok {
			return true
		}
		m[root.Val] = true
		root = root.Right
	}
	return false
}
