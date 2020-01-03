package queue

import "fmt"

type User struct {
	Name string
	Seq  int
}

func (u *User) String() string {
	return fmt.Sprintf("%s:%d", u.Name, u.Seq)
}

type LoginQueue struct {
	queue Queue
}

func (q *LoginQueue) Enqueue(user *User) {
	q.queue.Enqueue(user)
	user.Seq = q.queue.Size()
}

func (q *LoginQueue) Dequeue() *User {
	val := q.queue.Dequeue()
	q.UpdateSeq()
	return val.(*User)
}

func (q *LoginQueue) UpdateSeq() {
	i := 1
	for _, item := range q.queue.List() {
		item.(*User).Seq = i
		i++
	}
}

func (q *LoginQueue) Traversal() {
	users := make([]interface{}, 0)
	for _, item := range q.queue.List() {
		users = append(users, item.(*User))
	}
	fmt.Println(users...)
}
