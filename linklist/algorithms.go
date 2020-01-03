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
func Sort1(l1 *ListNode) *ListNode {
	if l1 == nil || l1.Next == nil {
		return l1
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
	l2 := (*LinkList)(fn(l1)).Reverse()
	cur1, cur2 := l1, (*ListNode)(l2)
	for cur1.Next != nil {
		next := cur2.Next
		cur2.Next = cur1.Next
		cur1.Next = cur2
		cur1 = cur2.Next
		cur2 = next
	}
	cur1.Next = cur2
	return l1
}
