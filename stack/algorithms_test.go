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

func TestMvBottomToTopSort(t *testing.T) {
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
	tests := []struct {
		name string
		args args
	}{
		{"123456", args{s: fn(6, 5, 4, 3, 2, 1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MvBottomToTopSort(tt.args.s)
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

func TestSortStack(t *testing.T) {
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
	tests := []struct {
		name string
		args args
	}{
		{"123456", args{s: fn(1, 3, 2, 5, 4, 6)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortStack(tt.args.s)
			tt.args.s.Traversal()
		})
	}
}

func TestIsPopSerial(t *testing.T) {
	type args struct {
		in  []int
		out []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{[]int{}, []int{}}, true},
		{"empty1", args{[]int{}, []int{1}}, false},
		{"empty2", args{[]int{1}, []int{}}, false},
		{"right", args{[]int{1, 2, 3, 4, 5}, []int{3, 2, 5, 4, 1}}, true},
		{"fault", args{[]int{1, 2, 3, 4, 5}, []int{3, 5, 2, 4, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPopSerial(tt.args.in, tt.args.out); got != tt.want {
				t.Errorf("IsPopSerial() = %v, want %v", got, tt.want)
			}
		})
	}
}
