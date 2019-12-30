package arrayStack

import "fmt"

type Stack struct {
	cap   int
	top   int
	elems []interface{}
}

func NewStack(cap int) *Stack {
	return &Stack{
		cap:   cap,
		top:   -1,
		elems: make([]interface{}, cap),
	}
}

func (s *Stack) Traversal() {
	for i := s.top; i > -1; i-- {
		fmt.Print(s.elems[i], " ")
	}
	fmt.Println()
}

func (s *Stack) Push(val interface{}) {
	if s.cap == s.top+1 {
		if s.cap < 1000 {
			s.cap *= 2
		} else {
			s.cap += 1000
		}
		elems := make([]interface{}, s.cap)
		copy(elems, s.elems)
		s.elems = elems
	}
	s.top++
	s.elems[s.top] = val
}

func (s *Stack) Pop() interface{} {
	if s.top == -1 {
		return nil
	}
	val := s.elems[s.top]
	s.top--
	if s.top < s.cap/2 {
		s.cap /= 2
		elems := make([]interface{}, s.cap)
		copy(elems, s.elems)
		s.elems = elems
	}
	return val
}
