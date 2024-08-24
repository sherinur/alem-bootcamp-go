package bootcamp

func (l *List) Sort(fn func(a *ListNode, b *ListNode) int) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	front, back := l.splitList(l.Head)

	frontList := &List{Head: front}
	backList := &List{Head: back}
	frontList.Sort(fn)
	backList.Sort(fn)

	l.Head = l.merge(frontList.Head, backList.Head, fn)
}

func (l *List) splitList(head *ListNode) (*ListNode, *ListNode) {
	slow := head
	fast := head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	}
	return head, slow
}

func (l *List) merge(front *ListNode, back *ListNode, fn func(a *ListNode, b *ListNode) int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for front != nil && back != nil {
		if fn(front, back) <= 0 {
			current.Next = front
			front = front.Next
		} else {
			current.Next = back
			back = back.Next
		}
		current = current.Next
	}
	if front != nil {
		current.Next = front
	}
	if back != nil {
		current.Next = back
	}
	return dummy.Next
}
