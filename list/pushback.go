package bootcamp

func (l *List) PushBack(v interface{}) {
	newNode := &ListNode{
		Next:  nil,
		Value: v,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
