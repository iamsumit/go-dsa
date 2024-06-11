package st

type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}

	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		return -1
	}

	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Data() []int {
	return s.data
}
