package linklist

import "errors"

type SingleLinkedNode struct {
	Val  int
	Next *SingleLinkedNode
}

type SingleLinkedList struct {
	length int
	head   *SingleLinkedNode
	tail   *SingleLinkedNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) IsEmpty() bool {
	return l.length == 0
}

//向链表尾插入节点
func (l *SingleLinkedList) AddNode(val int) {
	node := &SingleLinkedNode{Val: val}
	if l.tail == nil {
		l.tail = node
		l.head = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
	l.length++
}

//在链表中插入一个节点
func (l *SingleLinkedList) InsertNode(val, idx int) error {
	if idx < 0 || idx >= l.length {
		return errors.New("index out of range")
	}
	if l.IsEmpty() || idx == l.length {
		l.AddNode(val)
		return nil
	}
	node := &SingleLinkedNode{Val: val}
	if idx == 0 {
		node.Next = l.head
		l.head = node
	} else {
		p := l.head
		for i := 0; i <= idx; i++ {
			p = p.Next
		}
		node.Next = p.Next
		p.Next = node
	}

	l.length++
	return nil
}

func (l *SingleLinkedList) DeleteNode(idx int) error {
	if idx < 0 || idx >= l.length {
		return errors.New("index out of range")
	}
	p := l.head
	for i := 1; i < idx; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	if p.Next == nil {
		l.tail = p
	}
	l.length--
	return nil
}

func (l *SingleLinkedList) FindNode(idx int) (int, error) {
	if idx < 0 || idx >= l.length {
		return 0, errors.New("index out of range")
	}
	p := l.head
	for i := 0; i <= idx; i++ {
		p = p.Next
	}
	return p.Val, nil
}

func (l *SingleLinkedList) Traverse() []int {
	result := make([]int, 0, l.length)
	p := l.head
	for p != nil {
		result = append(result, p.Val)
	}
	return result
}
