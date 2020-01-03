package stack

type MinStack struct {
	elemStack Stack
	minStack  Stack
	cmp       func(interface{}, interface{}) int
}

func (ms *MinStack) Push(val interface{}) {
	ms.elemStack.Push(val)
	if ms.minStack.IsEmpty() || ms.cmp(val, ms.minStack.Top()) < 0 {
		ms.minStack.Push(val)
	}
}

func (ms *MinStack) Pop() interface{} {
	if ms.elemStack.IsEmpty() {
		return nil
	}
	val := ms.elemStack.Pop()
	if ms.elemStack.IsEmpty() || ms.cmp(ms.elemStack.Top(), ms.minStack.Top()) > 0 {
		ms.minStack.Pop()
	}
	return val
}

func (ms *MinStack) Min() interface{} {
	return ms.minStack.Top()
}
