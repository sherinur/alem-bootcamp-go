package bootcamp

func (l *List) Len() int {
	current := l.Head
	counter := 0

	for current != nil {
		counter++
		current = current.Next
	}

	return counter
}
