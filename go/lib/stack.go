package lib

type Stack struct {
	value   []interface{}
	start   int
	isEmpty bool
	end     int
}

func NewStack(value []interface{}) (*Stack, error) {

	if value == nil {
		return nil, StackInputIsNil
	}

	q := Stack{
		value:   value,
		isEmpty: len(value) == 0,
		start:   0,
		end:     len(value) - 1,
	}
	return &q, nil
}

func (s *Stack) Top() (interface{}, error) {
	if len(s.value) == 0 {
		return nil, StackIsEmpty
	}
	return s.value[s.end], nil
}

func (s *Stack) Push(val interface{}) {

	if s.end == len(s.value)-1 {
		s.value = append(s.value, val)
		s.end++
		s.isEmpty = false
		return
	}

	if s.isEmpty {
		s.value[s.end] = val
	} else {
		s.end++
		s.value[s.end] = val
	}

	s.isEmpty = false

	return
}

func (s *Stack) IsEmpty() bool {
	return s.isEmpty
}

func (s *Stack) Pop() {

	if s.end > 0 {
		s.end--
		return
	}

	s.isEmpty = true
	return
}
