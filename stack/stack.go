package stack

type Stack interface {
	IsEmpty() bool
	Top() interface{}
	Push(val interface{})
	Pop() interface{}
	Traversal()
}
