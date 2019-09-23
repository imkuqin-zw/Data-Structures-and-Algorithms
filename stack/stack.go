package stack

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

type Stack struct {
	top    *Node
	length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.Data
}

func (s *Stack) Push(val interface{}) {
	node := NewNode(val)
	node.Next = s.top
	s.top = node
	s.length++
}

func (s *Stack) Pop() {
	if s.length == 0 {
		return
	}
	node := s.top.Next
	s.top.Next = nil
	s.top = node
	s.length--
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
