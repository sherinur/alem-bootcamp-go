package bootcamp

func (q *Queue) Peek() interface{} {
	// Return the item at the front of the queue without removing it
	if len(q.Items) == 0 {
		return nil
	}
	return q.Items[0]
}
