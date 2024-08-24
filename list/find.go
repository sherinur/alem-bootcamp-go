package bootcamp

func (l *List) Find(fn func(v interface{}) bool) *ListNode {
	if l.Head == nil {
		return nil
	}

	current := l.Head

	for current != nil {
		if fn(current.Value) {
			return current
		}
		current = current.Next
	}
	return nil
}
