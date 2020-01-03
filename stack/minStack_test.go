package stack

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/arrayStack"
	"reflect"
	"testing"
)

func TestMinStack_Min(t *testing.T) {

	ms := &MinStack{
		elemStack: arrayStack.NewStack(),
		minStack:  arrayStack.NewStack(),
		cmp: func(i interface{}, i2 interface{}) int {
			return i.(int) - i2.(int)
		},
	}

	tests := []struct {
		name string
		fn   func()
		want interface{}
	}{
		{"", func() {}, nil},
		{"", func() { ms.Pop() }, nil},
		{"", func() { ms.Push(1) }, 1},
		{"", func() { ms.Push(2) }, 1},
		{"", func() { ms.Push(0) }, 0},
		{"", func() { ms.Pop() }, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			if got := ms.Min(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
