package listStack

import "fmt"

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

func (s *Stack) Traversal() {
	cur := s.top
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func (s *Stack) Push(val interface{}) {
	node := &Node{
		Data: val,
		Next: s.top,
	}
	s.top = node
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	node := s.top
	next := node.Next
	node.Next = nil
	s.top = next
	s.length--
	return node.Data
}

func (s *Stack) Top() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.Data
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
