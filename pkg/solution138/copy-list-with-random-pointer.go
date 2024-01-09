package solution138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 1. 创建原始节点的副本并插入到原始节点后
	for node := head; node != nil; node = node.Next.Next {
		newNode := &Node{
			Val:  node.Val,
			Next: node.Next,
		}
		node.Next = newNode
	}

	// 2. 设置副本节点的random指针
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	// 3. 把原始节点与复制的节点分离
	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = newNode.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}

	return newHead
}
