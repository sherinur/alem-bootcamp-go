package btree

func (b *BTree) PreOrderTraversal(fn func(n *BTreeNode)) {
	preOrderHelper(b.Root, fn)
}

func preOrderHelper(node *BTreeNode, fn func(n *BTreeNode)) {
	if node != nil {
		fn(node)
		preOrderHelper(node.Left, fn)
		preOrderHelper(node.Right, fn)
	}
}
