package btree

func (b *BTree) PostOrderTraversal(fn func(n *BTreeNode)) {
	postOrderHelper(b.Root, fn)
}

func postOrderHelper(node *BTreeNode, fn func(n *BTreeNode)) {
	if node != nil {
		postOrderHelper(node.Left, fn)
		postOrderHelper(node.Right, fn)
		fn(node)

	}
}
