package btree

func (b *BTree) DeleteNode(value int) {
	if b.Root == nil {
		return
	}
	b.Root = deleteNode(b.Root, value)
}

func deleteNode(node *BTreeNode, value int) *BTreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = deleteNode(node.Right, minNode.Value)
	}

	return node
}

func findMin(node *BTreeNode) *BTreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
