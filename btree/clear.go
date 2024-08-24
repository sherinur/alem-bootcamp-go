package btree

func (b *BTree) Clear() {
	b.Root = nil
}
