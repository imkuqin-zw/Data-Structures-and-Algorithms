package linklist

import (
	"fmt"
	"strings"
)

type CrossNode struct {
	Val  interface{}
	Next *CrossNode
	Down *CrossNode
}

func (c *CrossNode) String() string {
	cur := c
	str := strings.Builder{}
	for cur != nil {
		if cur.Down != nil {
			str.WriteString(fmt.Sprintf("%v,", cur.Val))
		} else {
			str.WriteString(fmt.Sprintf("%v", cur.Val))
		}
		cur = cur.Down
	}
	return str.String()
}

func CreateList(values [][]interface{}) *CrossNode {
	if len(values) == 0 {
		return nil
	}
	h := new(CrossNode)
	cur := h
	for _, vals := range values {
		inH := &CrossNode{Val: vals[0]}
		cur.Next = inH
		cur = cur.Next
		inCur := inH
		for i := 1; i < len(vals); i++ {
			inCur.Down = &CrossNode{Val: vals[i]}
			inCur = inCur.Down
		}
	}
	return h.Next
}

func Flatten(l *CrossNode, cmp func(val1, val2 interface{}) int) *CrossNode {
	if l == nil {
		return l
	}
	head := l
	for head.Next != nil {
		next := head.Next.Next
		head = MergeSortCrossLink(head, head.Next, cmp)
		head.Next = next
	}
	return head
}

func MergeSortCrossLink(l1 *CrossNode, l2 *CrossNode, cmp func(val1, val2 interface{}) int) *CrossNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l1.Next = nil
	l2.Next = nil
	var (
		head, pre, cur1, cur2 *CrossNode
	)
	if cmp(l1.Val, l2.Val) <= 0 {
		head, pre, cur1, cur2 = l1, l1, l1.Down, l2
	} else {
		head, pre, cur1, cur2 = l2, l2, l2.Down, l1
	}
	for cur1 != nil && cur2 != nil {
		if cmp(cur1.Val, cur2.Val) <= 0 {
			pre.Down = cur1
			pre = cur1
			cur1 = cur1.Down
		} else {
			pre.Down = cur2
			pre = cur2
			cur2 = cur2.Down
		}
	}
	if cur1 != nil {
		pre.Down = cur1
	}
	if cur2 != nil {
		pre.Down = cur2
	}
	return head
}
