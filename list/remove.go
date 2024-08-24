package bootcamp

// This method remove the given node from the doubly linked list.
func (l *List) Remove(n *ListNode) {
	if l.Head == nil {
		return
	}

	if n == l.Head {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next == n {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
