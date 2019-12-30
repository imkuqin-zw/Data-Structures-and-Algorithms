package arrayAqueue

import "fmt"

type Queue struct {
	cap   int
	rear  int
	elems []interface{}
}

func NewQueue(cap int) *Queue {
	return &Queue{
		cap:   cap,
		rear:  -1,
		elems: make([]interface{}, cap),
	}
}

func (s *Queue) Enqueue(val interface{}) {
	if s.rear+1 == s.cap {
		if s.cap < 1000 {
			s.cap *= 2
		} else {
			s.cap += 1000
		}
		elems := make([]interface{}, s.cap)
		copy(elems, s.elems)
		s.elems = elems
	}
	for i := s.rear; i > -1; i-- {
		s.elems[i+1] = s.elems[i]
	}
	s.rear++
	s.elems[0] = val
}

func (s *Queue) Dequeue() interface{} {
	if s.rear == -1 {
		return nil
	}
	val := s.elems[s.rear]
	s.rear--
	if s.rear < s.cap/2 {
		s.cap /= 2
		elems := make([]interface{}, s.cap)
		copy(elems, s.elems)
		s.elems = elems
	}
	return val
}

func (s *Queue) Traversal() {
	for i := s.rear; i > -1; i-- {
		fmt.Print(s.elems[i], " ")
	}
	fmt.Println()
}
