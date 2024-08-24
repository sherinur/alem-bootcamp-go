package bootcamp

import "bootcamp/btree"

func LevelsBtree(b *btree.BTree) int {
	if b.Root == nil {
		return 0
	}
	return calculateLevels(b.Root)
}

func calculateLevels(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}

	leftLevels := calculateLevels(node.Left)
	rightLevels := calculateLevels(node.Right)

	if leftLevels > rightLevels {
		return leftLevels + 1
	}
	return rightLevels + 1
}
