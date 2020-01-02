package stack

func MvBottomToTop(s Stack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Top()
	if !s.IsEmpty() {
		MvBottomToTop(s)
		top2 := s.Top()
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
