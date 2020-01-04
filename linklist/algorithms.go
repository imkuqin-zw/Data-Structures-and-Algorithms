package linklist

func Add(l1, l2 *ListNode) *ListNode {
	var l3 = new(ListNode)
	if l1 == nil && l2 == nil {
		return &ListNode{Val: 0}
	}
	cur, cur1, cur2 := l3, l1, l2
	c := 0
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val.(int) + cur2.Val.(int) + c
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		c = sum / 10
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	for cur1 != nil {
		sum := cur1.Val.(int) + c
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		c = sum / 10
		cur1 = cur1.Next
	}
	for cur2 != nil {
		sum := cur2.Val.(int) + c
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		c = sum / 10
		cur2 = cur2.Next
	}
	return l3.Next
}

/**
链表重新排序
1,2,3,4,5,6,7,8 => 1,8,2,7,3,6,4,5
*/
func Sort1(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	fn := func(l1 *ListNode) *ListNode {
		pre, slow, fast := l1, l1, l1
		for fast != nil && fast.Next != nil {
			pre = slow
			slow = slow.Next
			fast = fast.Next.Next
		}
		pre.Next = nil
		return slow
	}
	l2 := (*LinkList)(fn(l)).Reverse()
	cur1, cur2 := l, (*ListNode)(l2)
	for cur1.Next != nil {
		next := cur2.Next
		cur2.Next = cur1.Next
		cur1.Next = cur2
		cur1 = cur2.Next
		cur2 = next
	}
	cur1.Next = cur2
	return l
}

func FindLastK(l *ListNode, k int) interface{} {
	slow, fast := l, l
	i := 0
	for ; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if i < k {
		return nil
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

func IsLoop(l *ListNode) *ListNode {
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

func GetLoopNode(l *ListNode) *ListNode {
	meet := IsLoop(l)
	if meet == nil {
		return nil
	}
	head := l
	for head != meet {
		head = head.Next
		meet = meet.Next
	}
	return head
}

func SwapNeighboring(l **ListNode) {
	if l == nil || (*l) == nil || (*l).Next == nil {
		return
	}
	tmp := (*l).Next
	(*l).Next = tmp.Next
	tmp.Next = *l
	*l = tmp
	pre := (*l).Next
	cur := pre.Next
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
		pre = cur
		cur = pre.Next
	}
	return
}

func MergeSortLink(l1 *ListNode, l2 *ListNode, cmp func(val1, val2 interface{}) int) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var (
		head, pre, cur1, cur2 *ListNode
	)
	if cmp(l1.Val, l2.Val) <= 0 {
		head, pre, cur1, cur2 = l1, l1, l1.Next, l2
	} else {
		head, pre, cur1, cur2 = l2, l2, l2.Next, l1
	}
	for cur1 != nil && cur2 != nil {
		if cmp(cur1.Val, cur2.Val) <= 0 {
			pre.Next = cur1
			pre = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			pre = cur2
			cur2 = cur2.Next
		}
	}
	if cur1 != nil {
		pre.Next = cur1
	}
	if cur2 != nil {
		pre.Next = cur2
	}
	return head
}
