package arrayStack

import (
	"fmt"
)

type Stack struct {
	size  int
	elems []interface{}
}

func NewStack() *Stack {
	return &Stack{
		elems: make([]interface{}, 0),
	}
}

func (s *Stack) Traversal() {
	for i := s.size - 1; i >= 0; i-- {
		fmt.Print(s.elems[i], " ")
	}
	fmt.Println()
}

func (s *Stack) Push(val interface{}) {
	s.elems = append(s.elems, val)
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	s.size--
	val := s.elems[s.size]
	s.elems = s.elems[:s.size]
	return val
}

func (s *Stack) Top() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.elems[s.size-1]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
