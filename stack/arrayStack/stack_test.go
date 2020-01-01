package arrayStack

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	fn := func(vals ...interface{}) {
		s := NewStack()
		for _, val := range vals {
			s.Push(val)
		}
		s.Traversal()
	}
	fn(1, 2, 3, 4, 5, 6)
	fn(7, 8, 9, 10, 11, 12, 13)
}

func TestStack_Pop(t *testing.T) {
	fn := func(vals ...interface{}) {
		s := NewStack()
		for _, val := range vals {
			s.Push(val)
		}
		s.Traversal()
		for {
			val := s.Pop()
			if val == nil {
				break
			}
			fmt.Print(val, " ")
		}
		fmt.Println()
		s.Traversal()
	}
	fn(1, 2, 3, 4, 5, 6)
	//fn(7, 8, 9, 10, 11, 12, 13)
}
