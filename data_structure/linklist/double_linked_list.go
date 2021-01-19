package linklist

import "errors"

type DoubleLinkedNode struct {
	Val  int
	Next *DoubleLinkedNode
	Prev *DoubleLinkedNode
}

type DoubleLinkedList struct {
	length int
	head   *DoubleLinkedNode
	tail   *DoubleLinkedNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (l *DoubleLinkedList) IsEmpty() bool {
	return l.length == 0
}

//向链表尾插入节点
func (l *DoubleLinkedList) AddNode(val int) {
	node := &DoubleLinkedNode{Val: val}
	if l.tail == nil {
		l.tail = node
		l.head = node
	} else {
		l.tail.Next = node
		node.Prev = l.tail
		l.tail = node

	}
	l.length++
}

//在链表中插入一个节点
func (l *DoubleLinkedList) InsertNode(val, idx int) error {
	if idx < 0 || idx >= l.length {
		return errors.New("index out of range")
	}
	if l.IsEmpty() || idx == l.length {
		l.AddNode(val)
		return nil
	}
	node := &DoubleLinkedNode{Val: val}
	if idx == 0 {
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	} else {
		p := l.head
		for i := 0; i <= idx; i++ {
			p = p.Next
		}
		node.Prev = p
		node.Next = p.Next
		p.Next.Prev = node
		p.Next = node
	}

	l.length++
	return nil
}

func (l *DoubleLinkedList) DeleteNode(idx int) error {
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
	} else {
		p.Next.Prev = p
	}
	l.length--
	return nil
}

func (l *DoubleLinkedList) FindNode(idx int) (int, error) {
	if idx < 0 || idx >= l.length {
		return 0, errors.New("index out of range")
	}
	p := l.head
	for i := 0; i <= idx; i++ {
		p = p.Next
	}
	return p.Val, nil
}

func (l *DoubleLinkedList) Traverse() []int {
	result := make([]int, 0, l.length)
	p := l.head
	for p != nil {
		result = append(result, p.Val)
	}
	return result
}
