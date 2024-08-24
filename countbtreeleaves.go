package bootcamp

import "bootcamp/btree"

func CountBtreeLeaves(b *btree.BTree) int {
	if b.Root == nil {
		return 0
	}
	return countLeaves(b.Root)
}

func countLeaves(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return countLeaves(node.Left) + countLeaves(node.Right)
}
