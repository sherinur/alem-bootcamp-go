package bootcamp

import "bootcamp/btree"

func IsBalancedBtree(b *btree.BTree) bool {
	return IsBalancedBtreeNode(b.Root)
}

func IsBalancedBtreeNode(b *btree.BTreeNode) bool {
	if b == nil {
		return true
	}

	leftHeight := height(b.Left)
	rightHeight := height(b.Right)

	if (leftHeight - rightHeight) > 1 {
		return false
	}

	return IsBalancedBtreeNode(b.Left) && IsBalancedBtreeNode(b.Right)
}

func height(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
