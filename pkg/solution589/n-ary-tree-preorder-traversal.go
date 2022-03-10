package solution589

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	return preorderRecursive(root)
}

func preorderRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	for _, child := range root.Children {
		res = append(res, preorderRecursive(child)...)
	}
	return res
}
