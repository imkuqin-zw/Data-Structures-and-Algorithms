package queue

type MyCircularQueue struct {
	head, tail int
	vals       []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head: -1,
		tail: -1,
		vals: make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.vals) == 0 {
		return false
	}
	if this.tail == -1 {
		this.head = 0
		this.tail = 0
		this.vals[this.tail] = value
		return true
	}
	next := (this.tail + 1) % len(this.vals)
	if next != this.head {
		this.tail = next
		this.vals[this.tail] = value
		return true
	}
	return false
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.head == -1 {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	if this.head%len(this.vals) != this.tail {
		this.head = (this.head + 1) % len(this.vals)
		return true
	}
	return false
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.head == -1 {
		return -1
	}
	return this.vals[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.tail == -1 {
		return -1
	}
	return this.vals[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%len(this.vals) == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
