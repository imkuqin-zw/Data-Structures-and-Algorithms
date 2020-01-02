package stack

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/arrayStack"
	"testing"
)

func TestMvBottomToTop(t *testing.T) {
	type args struct {
		s Stack
	}

	fn := func(vals ...interface{}) Stack {
		s := &arrayStack.Stack{}
		for _, val := range vals {
			s.Push(val)
		}
		return s
	}
	fn(1, 2, 3, 4, 5, 6)
	tests := []struct {
		name string
		args args
	}{
		{"123456", args{s: fn(1, 2, 3, 4, 5, 6)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MvBottomToTop(tt.args.s)
			tt.args.s.Traversal()
		})
	}
}

func TestReverseStack(t *testing.T) {
	type args struct {
		s Stack
	}
	fn := func(vals ...interface{}) Stack {
		s := &arrayStack.Stack{}
		for _, val := range vals {
			s.Push(val)
		}
		return s
	}
	fn(1, 2, 3, 4, 5, 6)
	tests := []struct {
		name string
		args args
	}{
		{"123456", args{s: fn(1, 2, 3, 4, 5, 6)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseStack(tt.args.s)
			tt.args.s.Traversal()
		})
	}
}
