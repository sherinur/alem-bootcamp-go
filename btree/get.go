package btree

func (b *BTree) Get(value int) *BTreeNode {
	return getRecursive(b.Root, value)
}

func getRecursive(node *BTreeNode, value int) *BTreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		return getRecursive(node.Left, value)
	} else if value > node.Value {
		return getRecursive(node.Right, value)
	}
	return node
}
