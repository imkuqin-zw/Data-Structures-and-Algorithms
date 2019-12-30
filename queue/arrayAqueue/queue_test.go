package arrayAqueue

import (
	"fmt"
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	fn := func(vals ...interface{}) {
		s := NewQueue(5)
		fmt.Println(s.cap)
		for _, val := range vals {
			s.Enqueue(val)
			s.Traversal()
		}
		fmt.Println("----------------------", s.cap)
		for s.Dequeue() != nil {
			s.Traversal()
		}
		fmt.Println(s.cap)
	}
	fn(1, 2, 3, 4, 5, 6)
}

func TestQueue_Enqueue(t *testing.T) {
	fn := func(vals ...interface{}) {
		s := NewQueue(5)
		for _, val := range vals {
			s.Enqueue(val)
		}
		s.Traversal()
	}
	fn(1, 2, 3, 4, 5, 6)
	fn(7, 8, 9, 10, 11, 12, 13)
}
