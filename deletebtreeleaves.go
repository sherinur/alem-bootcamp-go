package bootcamp

import "bootcamp/btree"

func DeleteBtreeLeaves(b *btree.BTree) {
	if b.Root == nil {
		return
	}
	deleteLeaves(b.Root)
}

func deleteLeaves(node *btree.BTreeNode) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		if node.Parent != nil {
			if node.Parent.Left == node {
				node.Parent.Left = nil
			} else if node.Parent.Right == node {
				node.Parent.Right = nil
			}
		}

		node.Parent = nil
		return
	}

	deleteLeaves(node.Left)
	deleteLeaves(node.Right)
}
