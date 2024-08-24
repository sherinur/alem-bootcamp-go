package bootcamp

func (l *List) PushBefore(n *ListNode, v interface{}) {
	if l.Head == nil {
		return
	}

	newNode := &ListNode{
		Next:  nil,
		Value: v,
	}

	if l.Head == n {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	current := l.Head

	for current.Next != nil {
		if current.Next == n {
			newNode.Next = current.Next
			current.Next = newNode
			break
		}
		current = current.Next
	}
}
