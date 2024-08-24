package bootcamp

func (l *List) ForEach(fn func(n *ListNode)) {
	if l.Head == nil {
		return
	}

	current := l.Head

	for current != nil {
		fn(current)
		current = current.Next
	}
}
