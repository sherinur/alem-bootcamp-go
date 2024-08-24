package bootcamp

import "bootcamp/btree"

func CountBtreeNodes(b *btree.BTree) int {
	return countNodesRecursive(b.Root)
}

func countNodesRecursive(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countNodesRecursive(node.Left) + countNodesRecursive(node.Right)
}
