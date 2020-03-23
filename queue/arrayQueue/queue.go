package arrayQueue

import (
	"fmt"
)

type Queue struct {
	front int
	rear  int
	elems []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		elems: make([]interface{}, 0),
	}
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *Queue) Size() int {
	return q.rear - q.front
}

func (q *Queue) GetFront() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.elems[q.front]
}

func (q *Queue) Enqueue(val interface{}) {
	q.elems = append(q.elems, val)
	q.rear++
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	val := q.elems[q.front]
	q.elems = q.elems[1:]
	q.rear--
	return val
}

func (q *Queue) List() []interface{} {
	return q.elems
}

func (q *Queue) Traversal() {
	for _, item := range q.elems {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
