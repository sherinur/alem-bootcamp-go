package bootcamp

func (l *List) PushBackLists(lists ...*List) {
	for _, list := range lists {
		list.ForEach(func(n *ListNode) {
			l.PushBack(n.Value)
		})
	}
}
