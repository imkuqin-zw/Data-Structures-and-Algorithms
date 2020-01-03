package queue

type Queue interface {
	GetFront() interface{}
	Enqueue(val interface{})
	Dequeue() interface{}
	Size() int
	Traversal()
	List() []interface{}
}
