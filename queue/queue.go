package queue

type Queue interface {
	GetFront() interface{}
	Enqueue(val interface{})
	Dequeue() interface{}
	Traversal()
}
