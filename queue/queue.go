package bootcamp

type Queue struct {
	Items []interface{}
}

func NewQueue() *Queue {
	// Initialize and return a new instance of Queue
	return &Queue{
		Items: make([]interface{}, 0),
	}
}

func (q *Queue) Push(v interface{}) {
	// Add an item to the end of the queue
	q.Items = append(q.Items, v)
}

func (q *Queue) Pop() interface{} {
	// Remove and return the item at the front of the queue
	if len(q.Items) == 0 {
		return nil
	}
	removed := q.Items[0]
	q.Items = q.Items[1:]
	return removed
}
