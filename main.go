package main

import (
	"fmt"
)

// ListNode is the node of the linklist
type ListNode struct {
	Val  int
	Next *ListNode
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	{
		var l = 1
		{
			fmt.Println(l)
		}
	}
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i
		if j != 0 && i < m && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i != 0 && j < n && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			var leftMax int
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(leftMax)
			}
			var rightMin int
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				rightMin = min(nums1[i], nums2[j])
			}
			return float64(leftMax+rightMin) / 2
		}
	}
	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	left := (m + n + 1) / 2
	right := (n + m + 2) / 2
	return (getKth(nums1, 0, m-1, nums2, 0, n-1, left) + getKth(nums1, 0, m-1, nums2, 0, n-1, right)) / 2
}

func getKth(nums1 []int, mMin, mMax int, nums2 []int, nMin, nMax, k int) float64 {
	m := mMax - mMin + 1
	n := nMax - nMin + 1
	if m > n {
		return getKth(nums2, nMin, nMax, nums1, mMin, mMax, k)
	}
	if m == 0 {
		return float64(nums2[nMin+k-1])
	}
	if k == 1 {
		return float64(min(nums1[mMin], nums2[nMin]))
	}
	leftMin := mMin + min(m, k/2) - 1
	rightMin := nMin + min(n, k/2) - 1
	if nums1[leftMin] > nums2[rightMin] {
		return getKth(nums1, mMin, mMax, nums2, rightMin+1, nMax, k-min(n, k/2))
	}
	return getKth(nums1, leftMin+1, mMax, nums2, nMin, nMax, k-min(m, k/2))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	index := make(map[byte]int)
	maxLen := 0
	left, right := 0, 0
	for right < n {
		if j, ok := index[s[right]]; ok {
			left = max(left, j)
		}
		index[s[right]] = right + 1
		right++
		maxLen = max(right-left, maxLen)

	}
	return maxLen
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	p, q := head, head
	for i := 0; i < k; i++ {
		if p.Next == nil {
			i++
			k %= i
			i = 0
			p = head
		}
		p = p.Next
	}
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	h := q.Next
	q.Next, p.Next = nil, head
	return h
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var preHead *ListNode
	var result = head
	for i := 0; i < m-1 && head != nil; i++ {
		preHead = head
		head = head.Next
	}
	modifyTail := head
	var p *ListNode
	for i := 0; head != nil && i < n-m+1; i++ {
		next := head.Next
		head.Next = p
		p = head
		head = next
	}
	modifyTail.Next = head
	if preHead == nil {
		result = p
	} else {
		preHead.Next = p
	}
	return result
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	i, j := 0, 0
	cursorA, cursorB := headA, headB
	for cursorA != nil {
		i++
		cursorA = cursorA.Next
	}
	for cursorB != nil {
		j++
		cursorB = cursorB.Next
	}
	if i < j {
		for k := 0; k < j-i; k++ {
			headB = headB.Next
		}
	} else if i > j {
		for k := 0; k < i-j; k++ {
			headA = headA.Next
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != nil {
				if head == fast {
					return head
				}
				head = head.Next
				fast = fast.Next
			}
		}
	}
	return nil
}

func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = true
		head = head.Next
	}
	return false
}

func partition(head *ListNode, x int) *ListNode {
	p, q := &ListNode{}, &ListNode{}
	curP, curQ := p, q
	for head != nil {
		next := head.Next
		if head.Val < x {
			curP.Next = head
			curP = curP.Next
		} else {
			curQ.Next = head
			curQ = curQ.Next
		}
		head = next
	}
	curP.Next = q.Next
	curQ.Next = nil
	return p.Next
}

// Node is the one of the listStack
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hash := make(map[*Node]int)
	index := make([]*Node, 0)
	cur := head
	i := 0
	for cur != nil {
		node := &Node{Val: cur.Val}
		index = append(index, node)
		hash[cur] = i
		i++
		cur = cur.Next
	}
	index = append(index, nil)
	cur = head
	i = 0
	for cur != nil {
		index[i].Next = index[i+1]
		if cur.Random != nil {
			index[i].Random = index[hash[cur.Random]]
		}
		i++
		cur = cur.Next
	}
	return index[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}
	var newList []*ListNode
	if length%2 == 1 {
		newList = make([]*ListNode, 0, length/2+1)
		newList = append(newList, lists[length-1])
	} else {
		newList = make([]*ListNode, 0, length/2)
	}

	for i := 0; i < length-1; i += 2 {
		newList = append(newList, mergeTwoLists(lists[i], lists[i+1]))
	}
	return mergeKLists(newList)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			tmp := fast.Val
			fast = fast.Next.Next
			for fast != nil && fast.Val == tmp {
				fast = fast.Next
			}
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = nil
	return dummy.Next
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dic := make(map[int]*ListNode)
	total := 0
	var h = head

	for ; head != nil; head = head.Next {
		total += head.Val

		if total == 0 {
			h = head.Next
			continue
		}

		if _, ok := dic[total]; !ok {
			dic[total] = head
			continue
		}

		k := dic[total]
		t := total
		for m := k.Next; m != head; m = m.Next {
			t += m.Val
			delete(dic, t)
		}
		k.Next = head.Next
	}
	return h
}

func main() {
}
