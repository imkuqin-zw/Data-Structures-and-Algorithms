package BTree

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/queue/arrayQueue"
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/listStack"
)

type BNode struct {
	Data                   interface{}
	Parent, LChild, RChild *BNode
	Height                 int
	Npl                    int
	Color                  int
}

type BTree BNode

func CreateBTree(vals []interface{}, start, end int) *BTree {
	var root *BNode
	if start <= end {
		mid := (start + end + 1) / 2
		root = &BNode{
			Data:   vals[mid],
			LChild: (*BNode)(CreateBTree(vals, start, mid-1)),
			RChild: (*BNode)(CreateBTree(vals, mid+1, end)),
		}
	}
	return (*BTree)(root)
}

func (b *BNode) LevelTraversal() []interface{} {
	result := make([]interface{}, 0)
	if b == nil {
		return result
	}
	q := arrayQueue.NewQueue()
	q.Enqueue(b)
	for !q.IsEmpty() {
		node := q.Dequeue().(*BNode)
		result = append(result, node.Data)
		if node.LChild != nil {
			q.Enqueue(node.LChild)
		}
		if node.RChild != nil {
			q.Enqueue(node.RChild)
		}
	}
	return result
}

func (b *BNode) InOrderTraversal() []interface{} {
	result := make([]interface{}, 0)
	if b == nil {
		return result
	}
	s := listStack.NewStack()
	cur := b
	for {
		if cur != nil {
			s.Push(cur)
			cur = cur.LChild
		} else {
			if s.IsEmpty() {
				break
			}
			node := s.Pop().(*BNode)
			result = append(result, node.Data)
			cur = node.RChild
		}
	}
	return result
}

func (b *BNode) PreOrderTraversal() []interface{} {
	result := make([]interface{}, 0)
	if b == nil {
		return result
	}
	s := listStack.NewStack()
	cur := b
	for {
		if cur != nil {
			result = append(result, cur.Data)
			if cur.RChild != nil {
				s.Push(cur.RChild)
			}
			cur = cur.LChild
		} else {
			if s.IsEmpty() {
				break
			}
			cur = s.Pop().(*BNode)
		}
	}
	return result
}

func (b *BNode) PostOrderTraversal() []interface{} {
	result := make([]interface{}, 0)
	if b == nil {
		return result
	}
	s := listStack.NewStack()
	cur := b
	for {
		if cur != nil {
			l := cur.LChild
			cur.LChild = nil
			s.Push(cur)
			cur = l
		} else {
			if s.IsEmpty() {
				break
			}
			node := s.Top().(*BNode)
			cur = node.RChild
			if node.RChild != nil {
				node.RChild = nil
			} else {
				s.Pop()
				result = append(result, node.Data)
			}
		}
	}
	return result
}
