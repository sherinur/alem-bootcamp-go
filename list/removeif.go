package bootcamp

func (l *List) RemoveIf(fn func(n *ListNode) bool) {
	if l.Head == nil {
		return
	}

	current := l.Head

	for current != nil {
		if fn(current) {
			l.Remove(current)
		}

		current = current.Next
	}
}
