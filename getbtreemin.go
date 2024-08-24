package bootcamp

import "bootcamp/btree"

func GetMin(b *btree.BTree) *btree.BTreeNode {
	if b.Root == nil {
		return nil
	}
	current := b.Root
	for current.Left != nil {
		current = current.Left
	}
	return current
}
