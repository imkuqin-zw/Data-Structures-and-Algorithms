package listQueue

import "fmt"

type Node struct {
	Data interface{}
	next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

type Queue struct {
	rear   *Node
	front  *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) Front() interface{} {
	if q.rear == q.front {
		return nil
	}
	return q.front.Data
}

//队尾插入
func (q *Queue) Enqueue(val interface{}) {
	node := NewNode(val)
	if q.rear == nil {
		q.rear = node
		q.front = q.rear
	} else {
		q.rear.next = node
		q.rear = node
	}

	q.length++
}

//对头删除
func (q *Queue) Dequeue() interface{} {
	if q.front == nil {
		return nil
	}
	node := q.front
	q.front = node.next
	node.next = nil
	if q.front == nil {
		q.rear = nil
	}
	q.length--
	return node.Data
}

func (q *Queue) Empty() bool {
	return q.length == 0
}

func (q *Queue) Traversal() {
	cur := q.front
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.next
	}
	fmt.Println()
}

func (q *Queue) List() []interface{} {
	result := make([]interface{}, 0, q.length)
	cur := q.front
	for cur != nil {
		result = append(result, cur.Data)
		cur = cur.next
	}
	return result
}
