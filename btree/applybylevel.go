package btree

func (b *BTree) ApplyByLevel(fn func(node *BTreeNode, level int)) {
	if b.Root == nil {
		return
	}

	queue := []*BTreeNode{b.Root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[i]
			fn(node, level)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[levelSize:]
		level++
	}
}
