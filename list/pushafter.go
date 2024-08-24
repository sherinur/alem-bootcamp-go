package bootcamp

func (l *List) PushAfter(n *ListNode, v interface{}) {
	if l.Head == nil {
		return
	}

	newNode := &ListNode{
		Next:  nil,
		Value: v,
	}
	current := l.Head

	for current != nil {
		if current == n {
			newNode.Next = current.Next
			current.Next = newNode
		}
		current = current.Next
	}
}
