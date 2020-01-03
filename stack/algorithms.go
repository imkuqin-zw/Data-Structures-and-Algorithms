package stack

import "github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/arrayStack"

func MvBottomToTop(s Stack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Pop()
	if !s.IsEmpty() {
		MvBottomToTop(s)
		top2 := s.Pop()
		s.Push(top1)
		s.Push(top2)
	} else {
		s.Push(top1)
	}
}

func ReverseStack(s Stack) {
	if s.IsEmpty() {
		return
	}
	MvBottomToTop(s)
	top := s.Pop()
	ReverseStack(s)
	s.Push(top)
}

func MvBottomToTopSort(s Stack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Pop()
	if !s.IsEmpty() {
		MvBottomToTopSort(s)
		top2 := s.Top()
		if top2.(int) < top1.(int) {
			s.Pop()
			s.Push(top1)
			s.Push(top2)
			return
		}
	}
	s.Push(top1)
}

func SortStack(s Stack) {
	if s.IsEmpty() {
		return
	}
	MvBottomToTopSort(s)
	top := s.Pop()
	SortStack(s)
	s.Push(top)
}

func IsPopSerial(in, out []int) bool {
	inLen, outLen := len(in), len(out)
	if inLen != outLen {
		return false
	}
	s := arrayStack.NewStack()
	inIndex, outIndex := 0, 0
	for inIndex < inLen {
		s.Push(in[inIndex])
		inIndex++
		for !s.IsEmpty() && s.Top().(int) == out[outIndex] {
			s.Pop()
			outIndex++
		}
	}
	if !s.IsEmpty() || outIndex != outLen {
		return false
	}
	return true
}
