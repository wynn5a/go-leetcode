package solution

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoSortedLists Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists
func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// return another if one is nil
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	//choose smaller one as head
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		//append the rest of l1
		node.Next = l1
	}

	if l2 != nil {
		//append the rest of l2
		node.Next = l2
	}

	return head
}

func Run21() {
	fmt.Println("--- LeetCode NO.21 begin")
	rand.Seed(time.Now().UTC().UnixNano())
	var header1 = ListNode{
		Val:  1,
		Next: nil,
	}
	var header2 = ListNode{
		Val:  1,
		Next: nil,
	}
	sortedArray := make([]int, 10)
	rand.Seed(time.Now().UTC().UnixNano())
	min := 2
	max := 30
	for i := range sortedArray {
		sortedArray[i] = rand.Intn(max-min+1) + min
	}
	sort.Ints(sortedArray)
	for i := 0; i < 10; i++ {
		node := ListNode{
			Val:  sortedArray[i],
			Next: nil,
		}
		if i < 5 {
			appendTo(&node, &header1)
		} else {
			appendTo(&node, &header2)
		}

	}
	fmt.Println("input list 1     :", toString(&header1))
	fmt.Println("input list 2     :", toString(&header2))
	res := mergeTwoSortedLists(&header1, &header2)
	fmt.Println("answer           :", toString(res))
	fmt.Println("LeetCode NO.21 end")
}

func appendTo(node *ListNode, header *ListNode) {
	if header == nil || node == nil {
		return
	}
	for p := header; p != nil; p = p.Next {
		if p.Next == nil {
			p.Next = node
			break
		}
	}
}

func toString(node *ListNode) string {
	var result = ""
	for p := node; p != nil; p = p.Next {
		if p.Next == nil {
			result += strconv.Itoa(p.Val)
		} else {
			result += strconv.Itoa(p.Val) + " -> "
		}
	}
	return result
}
