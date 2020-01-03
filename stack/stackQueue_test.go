package stack

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/arrayStack"
	"reflect"
	"testing"
)

func TestStackQueue_Front(t *testing.T) {
	sq := Queue{
		elemStack:    arrayStack.NewStack(),
		dequeueStack: arrayStack.NewStack(),
	}
	tests := []struct {
		name string
		fn   func()
		want interface{}
	}{
		{"", func() {}, nil},
		{"", func() { sq.Enqueue(1) }, 1},
		{"", func() { sq.Enqueue(2) }, 1},
		{"", func() { sq.Enqueue(3) }, 1},
		{"", func() { sq.Dequeue() }, 2},
		{"", func() { sq.Dequeue() }, 3},
		{"", func() { sq.Dequeue() }, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			if got := sq.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Front() = %v, want %v", got, tt.want)
			}
		})
	}
}
