package bootcamp

type Stack struct {
	Items []interface{}
}

func NewStack() *Stack {
	// Initialize and return a new instance of Stack
	return &Stack{
		Items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(v interface{}) {
	s.Items = append(s.Items, v)
}

func (s *Stack) Pop() interface{} {
	// Remove and return the item at the top of the stack
	if len(s.Items) == 0 {
		return nil
	}
	lastIndex := len(s.Items) - 1
	v := s.Items[lastIndex]
	s.Items = s.Items[:lastIndex]
	return v
}
