package bootcamp

import "bootcamp/btree"

func GetMax(b *btree.BTree) *btree.BTreeNode {
	if b.Root == nil {
		return nil
	}
	current := b.Root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
