package bootcamp

func (q *Queue) Len() int {
	// Return the number of items in the queue
	return len(q.Items)
}
