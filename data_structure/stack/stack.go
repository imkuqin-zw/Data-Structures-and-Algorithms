package stack

type Stack struct {
	val []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.val = append(s.val, val)
}

func (s *Stack) Top() interface{} {
	if !s.Empty() {
		return s.val[s.Length()-1]
	}
	return nil
}

func (s *Stack) Pop() interface{} {
	if !s.Empty() {
		val := s.val[s.Length()-1]
		s.val = s.val[:s.Length()-1]
		return val
	} else {
		return nil
	}
}

func (s *Stack) Length() int {
	return len(s.val)
}

func (s *Stack) Empty() bool {
	return s.Length() > 0
}
