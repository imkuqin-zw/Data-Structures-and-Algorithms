package queue

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/queue/arrayAqueue"
	"testing"
)

func TestLoginQueue_Traversal(t *testing.T) {
	q := &LoginQueue{queue: arrayAqueue.NewQueue()}
	tests := []struct {
		fn func()
	}{
		{func() { q.Enqueue(&User{Name: "1"}) }},
		{func() { q.Enqueue(&User{Name: "2"}) }},
		{func() { q.Enqueue(&User{Name: "3"}) }},
		{func() { q.Dequeue() }},
		{func() { q.Dequeue() }},
	}
	for _, tt := range tests {
		tt.fn()
		q.Traversal()
	}
}
