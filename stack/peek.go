package bootcamp

func (s *Stack) Peek() interface{} {
	if len(s.Items) == 0 {
		return nil
	}
	// Return the top item of the stack without removing it
	return s.Items[len(s.Items)-1]
}
