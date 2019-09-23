/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-15 01:02:29
 * @LastEditTime: 2019-09-15 13:48:00
 * @LastEditors: Please set LastEditors
 */

package main

import (
	"container/list"
)

// // MinStack is wo
// type MinStack struct {
// 	data *Stack
// 	min  *Stack
// }

// // Push is
// func (m *MinStack) Push(val int) {
// 	m.data.Push(val)
// 	if m.min.length == 0 {
// 		m.min.Push(val)
// 	} else {
// 		m.min.Push(min(m.min.head.Val, val))
// 	}
// }

// // Pop is
// func (m *MinStack) Pop() (*ListNode, int) {

// 	return m.data.Pop(), m.min.Pop().Val
// }

// // Min is
// func (m *MinStack) Min() int {
// 	return m.min.head.Val
// }

// // Stack is
// type Stack struct {
// 	head   *ListNode
// 	top    *ListNode
// 	length int
// }

// // Push is
// func (s *Stack) Push(val int) {
// 	node := &ListNode{Val: val}
// 	node.Next = s.head
// 	s.head = node
// 	s.length++
// }

// // Pop is
// func (s *Stack) Pop() int {
// 	if s.length == 0 {
// 		return 0
// 	}
// 	result := s.head
// 	s.head = result.Next
// 	result.Next = nil
// 	return result.Val
// }

// // QNode is
// type QNode struct {
// 	Next *QNode
// 	Pre  *QNode
// 	Val  int
// }

// // Queue is
// type Queue struct {
// 	head   *QNode
// 	last   *QNode
// 	length int
// }

// // Push is
// func (s *Queue) Push(val int) {
// 	node := &QNode{Val: val, Next: s.head, Pre: nil}
// 	s.head.Pre = node
// 	s.head = node
// 	s.length++
// }

// // Pop is
// func (s *Queue) Pop() *QNode {
// 	if s.length == 0 {
// 		return nil
// 	}
// 	result := s.last
// 	s.last = s.last.Pre
// 	return result
// }

// // CheckIsValidOrder is
// func CheckIsValidOrder(o, o2 *Queue) bool {

// 	s := &Stack{}
// 	for i := 0; i < o.length; i++ {
// 		s.Push(o2.Pop().Val)
// 		for s.length > 0 && o.last.Val == s.head.Val {
// 			s.Pop()
// 			o.Pop()
// 		}
// 	}
// 	if s.length > 0 {
// 		return false
// 	}
// 	return true
// }

// // MyStack is
// type MyStack struct {
// 	data *list.List
// }

// /** Initialize your data structure here. */
// func Constructor() MyStack {
// 	return MyStack{data: list.New()}
// }

// /** Push element x onto stack. */
// func (this *MyStack) Push(x int) {
// 	tmp := list.New()
// 	tmp.PushFront(x)
// 	for this.data.Len() != 0 {
// 		tmp.PushFront(this.data.Back().Value)
// 		this.data.Remove(this.data.Back())
// 	}
// 	for tmp.Len() > 0 {
// 		this.data.PushFront(tmp.Back().Value)
// 		tmp.Remove(this.data.Back())
// 	}
// }

// /** Removes the element on top of the stack and returns that element. */
// func (this *MyStack) Pop() int {
// 	e := this.data.Back()
// 	this.data.Remove(e)
// 	return e.Value.(int)
// }

// /** Get the top element. */
// func (this *MyStack) Top() int {
// 	return this.data.Back().Value.(int)
// }

// /** Returns whether the stack is empty. */
// func (this *MyStack) Empty() bool {
// 	return this.data.Len() == 0
// }

// type MyQueue struct {
// 	data *list.List
// }

// /** Initialize your data structure here. */
// func Constructor() MyQueue {
// 	return MyQueue{data: list.New()}
// }

// /** Push element x to the back of queue. */
// func (this *MyQueue) Push(x int) {
// 	tmp := list.New()
// 	for this.data.Len() != 0 {
// 		e := this.data.Front()
// 		tmp.PushFront(e.Value)
// 		this.data.Remove(e)
// 	}
// 	tmp.PushFront(x)
// 	for tmp.Len() > 0 {
// 		e := tmp.Front()
// 		this.data.PushFront(e.Value)
// 		tmp.Remove(e)
// 	}
// }

// /** Removes the element from in front of queue and returns that element. */
// func (this *MyQueue) Pop() int {
// 	e := this.data.Front()
// 	this.data.Remove(e)
// 	return e.Value.(int)
// }

// /** Get the front element. */
// func (this *MyQueue) Peek() int {
// 	return this.data.Front().Value.(int)
// }

// /** Returns whether the queue is empty. */
// func (this *MyQueue) Empty() bool {
// 	return this.data.Len() == 0
// }

// // Stack is
// type Stack struct {
// 	head   *ListNode
// 	top    *ListNode
// 	length int
// }

// // Push is
// func (s *Stack) Push(val int) {
// 	node := &ListNode{Val: val}
// 	node.Next = s.head
// 	s.head = node
// 	s.length++
// }

// // Pop is
// func (s *Stack) Pop() *ListNode {
// 	if s.length == 0 {
// 		return nil
// 	}
// 	result := s.head
// 	s.head = result.Next
// 	return result
// }

// func getNum(s string) int {
// 	result := 0
// 	n := len(s)
// 	for i:=0; i < n; i++ {
// 		result = result * 10 + s[i] - '0'
// 	}
// 	return result
// }

// StackList is
type StackList struct {
	data *list.List
}

// Push is
func (s *StackList) Push(val interface{}) {
	s.data.PushFront(val)
}

// Pop is
func (s *StackList) Pop() interface{} {
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

func (s *StackList) Top() interface{} {
	return s.data.Front().Value
}

// Length is
func (s *StackList) Length() int {
	return s.data.Len()
}

func calculate(s string) int {
	option := map[byte]map[byte]byte{
		'+': {'+': '>', '-': '>', '*': '<', '/': '<', '^': '<', '!': '<', '(': '<', ')': '>', '#': '>'},
		'-': {'+': '>', '-': '>', '*': '<', '/': '<', '^': '<', '!': '<', '(': '<', ')': '>', '#': '>'},
		'*': {'+': '>', '-': '>', '*': '>', '/': '>', '^': '<', '!': '<', '(': '<', ')': '>', '#': '>'},
		'/': {'+': '>', '-': '>', '*': '>', '/': '>', '^': '<', '!': '<', '(': '<', ')': '>', '#': '>'},
		'^': {'+': '>', '-': '>', '*': '>', '/': '>', '^': '<', '!': '<', '(': '<', ')': '>', '#': '>'},
		'!': {'+': '>', '-': '>', '*': '>', '/': '>', '^': '<', '!': '<', '(': '<', ')': '>', '#': '>'},
		'(': {'+': '<', '-': '<', '*': '<', '/': '<', '^': '<', '!': '<', '(': '<', ')': '=', '#': ' '},
		')': {'+': ' ', '-': ' ', '*': ' ', '/': ' ', '^': ' ', '!': ' ', '(': ' ', ')': ' ', '#': ' '},
		'#': {'+': '<', '-': '<', '*': '<', '/': '<', '^': '<', '!': '<', '(': '<', ')': ' ', '#': '='},
	}

	numS := &StackList{data: list.New()}
	optionS := &StackList{data: list.New()}
	length := len(s)
	optionS.Push(byte('#'))
	i := 0
	for i < length {
		num := 0
		haveNum := '0' <= s[i] && s[i] <= '9'
		for ; i < length && '0' <= s[i] && s[i] <= '9'; i++ {
			num = num*10 + int(s[i]-'0')
		}
		if haveNum {
			numS.Push(num)
		}
		if i >= length {
			break
		} else if s[i] == ' ' {
			i++
			continue
		}

		switch option[optionS.Top().(byte)][s[i]] {
		case '<':
			optionS.Push(s[i])
			i++
		case '=':
			optionS.Pop()
			i++
		case '>':
			computer(numS, optionS)
		}
	}

	for optionS.Length() > 0 {
		computer(numS, optionS)
	}
	return numS.Pop().(int)
}

func computer(num, opt *StackList) {
	option := opt.Pop().(byte)
	if option == '!' {
		num1 := num.Pop().(int)
		num.Push(factorial(num1))
	} else if option != '#' {
		num1, num2 := num.Pop(), num.Pop()
		switch option {
		case '+':
			num.Push(num1.(int) + num2.(int))
		case '-':
			num.Push(num2.(int) - num1.(int))
		case '*':
			num.Push(num1.(int) * num2.(int))
		case '/':
			num.Push(num1.(int) / num2.(int))
		case '^':
			num.Push(power(num2.(int), num1.(int)))
		}
	}
}

func power(num, count int) int {
	result := 1
	for count > 0 {
		result *= num
		count--
	}
	return result
}

func factorial(num int) int {
	if num == 0 {
		return 0
	}
	result := 1
	for num > 1 {
		result *= num
		num--
	}
	return result
}
