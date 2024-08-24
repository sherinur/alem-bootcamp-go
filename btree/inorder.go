package btree

func (b *BTree) InOrderTraversal(fn func(n *BTreeNode)) {
	inOrderRecursive(b.Root, fn)
}

func inOrderRecursive(node *BTreeNode, fn func(n *BTreeNode)) {
	if node != nil {
		inOrderRecursive(node.Left, fn)
		fn(node)
		inOrderRecursive(node.Right, fn)
	}
}
