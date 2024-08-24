package btree

func (b *BTree) Clone() *BTree {
	if b.Root == nil {
		return &BTree{}
	}

	var cloneNode func(node *BTreeNode, parent *BTreeNode) *BTreeNode
	cloneNode = func(node *BTreeNode, parent *BTreeNode) *BTreeNode {
		if node == nil {
			return nil
		}
		cloned := &BTreeNode{
			Parent: parent,
			Value:  node.Value,
		}
		cloned.Left = cloneNode(node.Left, cloned)
		cloned.Right = cloneNode(node.Right, cloned)
		return cloned
	}

	clonedRoot := cloneNode(b.Root, nil)

	return &BTree{Root: clonedRoot}
}
