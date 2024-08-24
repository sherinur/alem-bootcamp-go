package btree

type BTree struct {
	Root *BTreeNode
}

type BTreeNode struct {
	Parent      *BTreeNode
	Left, Right *BTreeNode
	Value       int
}

func NewBTree() *BTree {
	return &BTree{
		Root: nil,
	}
}

func (b *BTree) ReplaceOrInsert(v int) {
	newNode := &BTreeNode{Value: v}
	if b.Root == nil {
		b.Root = newNode
		return
	}

	current := b.Root

	for {
		if v < current.Value {
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				return
			}
			current = current.Left
		} else if v > current.Value {
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				return
			}
			current = current.Right
		} else {
			return
		}
	}
}
