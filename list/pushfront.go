package bootcamp

func (l *List) PushFront(v interface{}) {
	newNode := &ListNode{
		Next:  nil,
		Value: v,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
