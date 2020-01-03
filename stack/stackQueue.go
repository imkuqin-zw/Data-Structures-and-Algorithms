package stack

type Queue struct {
	elemStack    Stack
	dequeueStack Stack
}

func (sq *Queue) Enqueue(val interface{}) {
	sq.elemStack.Push(val)
}

func (sq *Queue) Dequeue() interface{} {
	if sq.dequeueStack.IsEmpty() {
		for !sq.elemStack.IsEmpty() {
			sq.dequeueStack.Push(sq.elemStack.Pop())
		}
	}
	return sq.dequeueStack.Pop()
}

func (sq *Queue) Front() interface{} {
	if sq.dequeueStack.IsEmpty() {
		for !sq.elemStack.IsEmpty() {
			sq.dequeueStack.Push(sq.elemStack.Pop())
		}
	}
	return sq.dequeueStack.Top()
}
