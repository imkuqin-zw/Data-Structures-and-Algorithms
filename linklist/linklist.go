package linklist

import (
	"fmt"
	"strings"
)

// ListNode is the node of the linklist
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// LinkList is headless linked list
type LinkList ListNode

func (l *LinkList) InsertAt(index int, val interface{}) *LinkList {
	node := &ListNode{Val: val}
	if l == nil {
		return (*LinkList)(node)
	}
	if index == 0 {
		node.Next = (*ListNode)(l)
		return (*LinkList)(node)
	}
	var cur = (*ListNode)(l)
	for i := 1; i < index && cur.Next != nil; i++ {
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = node
	return l
}

func (l *LinkList) InsertFront(val interface{}) *LinkList {
	node := &ListNode{Val: val}
	if l == nil {
		return (*LinkList)(node)
	}
	node.Next = (*ListNode)(l)
	return (*LinkList)(node)
}

func (l *LinkList) InsertBack(val interface{}) *LinkList {
	node := &ListNode{Val: val}
	if l == nil {
		return (*LinkList)(node)
	}
	cur := (*ListNode)(l)
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	return l
}

func (l *LinkList) DeleteAt(index int) (*LinkList, interface{}) {
	if l == nil {
		return nil, nil
	}
	if index == 0 {
		h := (*LinkList)(l.Next)
		l.Next = nil
		return h, l.Val
	}
	cur := (*ListNode)(l)
	for i := 1; i < index && cur.Next != nil; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = nil
		return l, next.Val
	}
	return l, nil
}

func (l *LinkList) Remove(val interface{}) *LinkList {
	var cur = (*ListNode)(l)
	if cur.Val == val {
		h := (*LinkList)(cur.Next)
		cur.Next = nil
		return h
	}
	for cur.Next != nil && cur.Next.Val != val {
		cur = cur.Next
	}
	if cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = nil
	}
	return l
}

func (l *LinkList) Reverse() *LinkList {
	var (
		pre *ListNode
		cur = (*ListNode)(l)
		h   = cur
	)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		h = cur
		cur = next
	}
	return (*LinkList)(h)
}

func (l *LinkList) String() string {
	if l == nil {
		return ""
	}
	var s strings.Builder
	cur := (*ListNode)(l)
	for cur != nil {
		if cur.Next != nil {
			s.WriteString(fmt.Sprintf("%v,", cur.Val))
		} else {
			s.WriteString(fmt.Sprintf("%v", cur.Val))
		}
		cur = cur.Next
	}
	return s.String()
}

func (l *LinkList) Traverse() {
	cur := (*ListNode)(l)
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func (l *LinkList) Size() int {
	var size int
	cur := (*ListNode)(l)
	for cur != nil {
		size++
		cur = cur.Next
	}
	return size
}

// LinkListH is headed linked list
type LinkListH struct {
	size int
	Next *ListNode
}

func (l *LinkListH) InsertAt(index int, val interface{}) {
	node := &ListNode{Val: val}
	if index == 0 || l.Next == nil {
		node.Next = l.Next
		l.Next = node
		l.size++
		return
	}
	cur := l.Next
	for i := 1; i < index && cur.Next != nil; i++ {
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = node
	l.size++
	return
}

func (l *LinkListH) InsertFront(val interface{}) {
	node := &ListNode{Val: val, Next: l.Next}
	l.Next = node
	l.size++
}

func (l *LinkListH) InsertBack(val interface{}) {
	node := &ListNode{Val: val}
	if l.Next == nil {
		l.Next = node
	}
	cur := l.Next
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	l.size++
}

func (l *LinkListH) DeleteAt(index int) interface{} {
	if l.size == 0 {
		return nil
	}
	if index == 0 {
		node := l.Next
		l.Next = node.Next
		node.Next = nil
		l.size--
		return node.Val
	}
	cur := l.Next
	for i := 1; i < index && cur.Next != nil; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		node := cur.Next
		cur.Next = node.Next
		node.Next = nil
		l.size--
		return node.Val
	}
	return nil
}

func (l *LinkListH) Remove(val interface{}) {
	if l.Next != nil && l.Next.Val == val {
		next := l.Next
		l.Next = next.Next
		next.Next = nil
		l.size--
		return
	}
	cur := l.Next
	for ; cur.Next != nil && cur.Next.Val != val; cur = cur.Next {
	}
	if cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = nil
		l.size--
	}
}

func (l *LinkListH) Reverse() {
	var (
		pre *ListNode
		cur = l.Next
	)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		l.Next = cur
		cur = next
	}
}

func (l *LinkListH) Traverse() {
	cur := l.Next
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func (l *LinkListH) Size() int {
	return l.size
}

func (l *LinkListH) String() string {
	var s strings.Builder
	cur := l.Next
	for cur != nil {
		if cur.Next != nil {
			s.WriteString(fmt.Sprintf("%v,", cur.Val))
		} else {
			s.WriteString(fmt.Sprintf("%v", cur.Val))
		}
		cur = cur.Next
	}
	return s.String()
}

//func RemoveDup(l *ListNode) {
//	if l == nil || l.Next == nil {
//		return
//	}
//
//	for outCur := l.Next; outCur != nil; outCur = outCur.Next {
//		for pre, cur := outCur, outCur.Next; cur != nil; cur = cur.Next {
//			if outCur.Val == cur.Val {
//				pre.Next = cur.Next
//			} else {
//				pre = cur
//			}
//		}
//	}
//	return
//}
//
//func Add(l1, l2 *ListNode) *ListNode {
//	l3 := new(ListNode)
//	if l1 == nil || l2 == nil {
//		l3.Next = &ListNode{Val: 0}
//		return l3
//	}
//	cur, cur1, cur2 := l3, l1.Next, l2.Next
//	c := 0
//	for cur1 != nil && cur2 != nil {
//		sum := cur1.Val.(int) + cur2.Val.(int) + c
//		cur.Next = &ListNode{Val: sum % 10}
//		cur = cur.Next
//		c = sum / 10
//		cur1 = cur1.Next
//		cur2 = cur2.Next
//	}
//	for cur1 != nil {
//		sum := cur1.Val + c
//		cur.Next = &ListNode{Val: sum % 10}
//		cur = cur.Next
//		c = sum / 10
//		cur1 = cur1.Next
//	}
//	for cur2 != nil {
//		sum := cur2.Val + c
//		cur.Next = &ListNode{Val: sum % 10}
//		cur = cur.Next
//		c = sum / 10
//		cur2 = cur2.Next
//	}
//	if l3.Next == nil {
//		l3.Next = &ListNode{Val: 0}
//	}
//	return l3
//}
