package queue

type Queue struct {
	val []interface{}
}

func (q *Queue) Push(val interface{}) {
	q.val = append(q.val, val)
}

func (q *Queue) Top() interface{} {
	if !q.Empty() {
		return q.val[0]
	}
	return nil
}

func (q *Queue) Pop() interface{} {
	if !q.Empty() {
		val := q.val[0]
		q.val = q.val[1:]
		return val
	} else {
		return nil
	}
}

func (q *Queue) Length() int {
	return len(q.val)
}

func (q *Queue) Empty() bool {
	return q.Length() > 0
}
