package bootcamp

func (l *List) Reverse() {
	var prev *ListNode
	var next *ListNode
	current := l.Head

	l.Tail = l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
