package linklisk

import (
	"fmt"
	"strings"
)

// ListNode is the node of the linklist
type ListNode struct {
	Val  int
	Next *ListNode
}

func ToString(h *ListNode) string {
	var result strings.Builder
	var cur = h.Next
	for cur != nil {
		result.WriteByte(byte(cur.Val + 48))
		cur = cur.Next
	}
	return result.String()
}

func Traverse(h *ListNode) {
	cur := h.Next
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func HeadInsert(l *ListNode, val int) {
	if l == nil {
		return
	}
	n := &ListNode{
		Val:  val,
		Next: nil,
	}
	n.Next = l.Next
	l.Next = n
}

func TailInsert(l *ListNode, val int) {
	if l == nil {
		return
	}
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	n := &ListNode{
		Val:  val,
		Next: nil,
	}
	cur.Next = n
}

// have head
func Reverse(l *ListNode) {
	if l == nil || l.Next == nil {
		return
	}
	cur := l.Next.Next
	l.Next.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = l.Next
		l.Next = cur
		cur = next
	}
}

func RemoveDup(l *ListNode) {
	if l == nil || l.Next == nil {
		return
	}

	for outCur := l.Next; outCur != nil; outCur = outCur.Next {
		for pre, cur := outCur, outCur.Next; cur != nil; cur = cur.Next {
			if outCur.Val == cur.Val {
				pre.Next = cur.Next
			} else {
				pre = cur
			}
		}
	}
	return
}

func Add(l1, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	if l1 == nil || l2 == nil {
		l3.Next = &ListNode{Val: 0}
		return l3
	}
	cur, cur1, cur2 := l3, l1.Next, l2.Next
	c := 0
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val + cur2.Val + c
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		c = sum / 10
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	for cur1 != nil {
		sum := cur1.Val + c
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		c = sum / 10
		cur1 = cur1.Next
	}
	for cur2 != nil {
		sum := cur2.Val + c
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		c = sum / 10
		cur2 = cur2.Next
	}
	if l3.Next == nil {
		l3.Next = &ListNode{Val: 0}
	}
	return l3
}

//no head
func ReverseNoHead(h *ListNode) {
	if h == nil || h.Next == nil {
		return
	}
	var pre, next *ListNode
	for h != nil {
		next = h.Next
		h.Next = pre
		pre = h
		h = next
	}
}

func HeadInsertNoHead(h *ListNode, val int) {
	p := &ListNode{Val: val, Next: h}
	h = p
}

func TailInsertNoHead(h *ListNode, val int) {
	p := &ListNode{Val: val, Next: h}
	if h == nil {
		h = p
	}
	cur := h
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = p
	return
}

func TraverseNoHead(h *ListNode) {
	cur := h
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}
