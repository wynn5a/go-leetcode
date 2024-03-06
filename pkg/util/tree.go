package util

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// Array2TreeNode construct binary tree from string
// input [1,2,3,null,4]
// output tree with root 1
func Array2TreeNode(values string) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	ints := values[1 : len(values)-1]

	var nodes []*TreeNode
	for _, v := range strings.Split(ints, ",") {
		if v == "null" {
			nodes = append(nodes, nil)
		} else {
			val, _ := strconv.Atoi(v)
			nodes = append(nodes, &TreeNode{Val: val})
		}
	}

	i := 0
	j := 0
	for i < len(nodes) {
		if nodes[i] == nil {
			i++
			continue
		}
		fill(nodes[i], nodes, (j*2)+1)
		i++
		j++
	}

	return nodes[0]
}

func fill(node *TreeNode, nodes []*TreeNode, i int) {
	if node == nil {
		return
	}
	if i < len(nodes) {
		node.Left = nodes[i]
	}
	if i+1 < len(nodes) {
		node.Right = nodes[i+1]
	}
}

func IntArray2TreeNode(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(input))
	for i, v := range input {
		nodes[i] = &TreeNode{Val: v}
	}

	i := 0
	j := 0
	for i < len(nodes) {
		if nodes[i] == nil {
			i++
			continue
		}
		fill(nodes[i], nodes, (j*2)+1)
		i++
		j++
	}

	return nodes[0]
}

func Tree2IntArray(result *TreeNode) []int {
	var resultArray []int
	if result == nil {
		return []int{}
	}

	queue := []*TreeNode{result}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			resultArray = append(resultArray, -1)
			continue
		}
		resultArray = append(resultArray, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	for i := len(resultArray) - 1; i >= 0; i-- {
		if resultArray[i] != -1 {
			resultArray = resultArray[:i+1]
			break
		}
	}

	return resultArray
}
