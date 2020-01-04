package linklist

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	fn := func(a []int, b []int) args {
		arg := args{}
		l1 := new(LinkListH)
		for _, val := range a {
			l1.InsertBack(val)
		}
		arg.l1 = l1.Next
		l2 := new(LinkListH)
		for _, val := range b {
			l2.InsertBack(val)
		}
		arg.l2 = l2.Next
		return arg
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"", fn([]int{}, []int{}), "0"},
		{"", fn([]int{}, []int{1}), "1"},
		{"", fn([]int{1}, []int{}), "1"},
		{"", fn([]int{1, 2, 3}, []int{3, 2, 1}), "4,4,4"},
		{"", fn([]int{1, 9, 3}, []int{3, 2, 1}), "4,1,5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (*LinkList)(Add(tt.args.l1, tt.args.l2)).String(); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort1(t *testing.T) {
	type args struct {
		l1 *ListNode
	}
	fn := func(vals ...interface{}) args {
		l := new(LinkListH)
		for _, val := range vals {
			l.InsertBack(val)
		}
		return args{l.Next}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", fn(), ""},
		{"", fn(1), "1"},
		{"", fn(1, 2), "1,2"},
		{"", fn(1, 2, 3), "1,3,2"},
		{"", fn(1, 2, 3, 4, 5, 6, 7, 8), "1,8,2,7,3,6,4,5"},
		{"", fn(1, 2, 3, 4, 5, 6, 7), "1,7,2,6,3,5,4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (*LinkList)(Sort1(tt.args.l1)).String(); got != tt.want {
				t.Errorf("Sort1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLastK(t *testing.T) {
	type args struct {
		l *ListNode
		k int
	}
	fn := func(vals []interface{}, k int) args {
		l := new(LinkListH)
		for _, val := range vals {
			l.InsertBack(val)
		}
		return args{l.Next, k}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"", fn([]interface{}{}, 1), nil},
		{"", fn([]interface{}{1}, 2), nil},
		{"", fn([]interface{}{1}, 1), 1},
		{"", fn([]interface{}{1, 2, 3, 4}, 2), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLastK(tt.args.l, tt.args.k); got != tt.want {
				t.Errorf("FindLastK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLoop(t *testing.T) {
	type args struct {
		l *ListNode
	}
	fn := func(vals []int, loop int) args {
		l := new(LinkListH)
		for _, val := range vals {
			l.InsertBack(val)
		}
		if length := len(vals); length != 0 && loop < length {
			l.GetNodeByIndex(length - 1).Next = l.GetNodeByIndex(loop)
		}
		return args{l.Next}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", fn([]int{}, 0), ""},
		{"", fn([]int{0}, 1), ""},
		{"", fn([]int{0}, 0), "0"},
		{"", fn([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0), "0"},
		{"", fn([]int{0, 1, 2, 3, 4, 5, 6, 7}, 1), "7"},
		{"", fn([]int{0, 1, 2, 3, 4, 5, 6, 7}, 7), "7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLoop(tt.args.l); got.String() != tt.want {
				t.Errorf("IsLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLoopNode(t *testing.T) {
	type args struct {
		l *ListNode
	}
	fn := func(vals []int, loop int) args {
		l := new(LinkListH)
		for _, val := range vals {
			l.InsertBack(val)
		}
		if length := len(vals); length != 0 && loop < length {
			l.GetNodeByIndex(length - 1).Next = l.GetNodeByIndex(loop)
		}
		return args{l.Next}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", fn([]int{}, 0), ""},
		{"", fn([]int{0}, 1), ""},
		{"", fn([]int{0}, 0), "0"},
		{"", fn([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0), "0"},
		{"", fn([]int{0, 1, 2, 3, 4, 5, 6, 7}, 1), "1"},
		{"", fn([]int{0, 1, 2, 3, 4, 5, 6, 7}, 7), "7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLoopNode(tt.args.l); got.String() != tt.want {
				t.Errorf("GetLoopNode() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}

func TestSwapNeighboring1(t *testing.T) {
	type args struct {
		l *ListNode
	}
	fn := func(vals ...interface{}) args {
		l := new(LinkListH)
		for _, val := range vals {
			l.InsertBack(val)
		}
		return args{l.Next}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", fn(), ""},
		{"", fn(1), "1"},
		{"", fn(1, 2), "2,1"},
		{"", fn(1, 2, 3, 4, 5, 6, 7, 8), "2,1,4,3,6,5,8,7"},
		{"", fn(1, 2, 3, 4, 5, 6, 7), "2,1,4,3,6,5,7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SwapNeighboring(&tt.args.l)
			got := (*LinkList)(tt.args.l).String()
			if got != tt.want {
				t.Errorf("SwapNeighboring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortLink(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	cmp := func(val1, val2 interface{}) int {
		return val1.(int) - val2.(int)
	}
	fn := func(val1, val2 []int) args {
		l1 := new(LinkListH)
		for _, val := range val1 {
			l1.InsertBack(val)
		}
		l2 := new(LinkListH)
		for _, val := range val2 {
			l2.InsertBack(val)
		}
		return args{l1: l1.Next, l2: l2.Next}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", fn([]int{}, []int{}), ""},
		{"", fn([]int{1}, []int{}), "1"},
		{"", fn([]int{}, []int{1}), "1"},
		{"", fn([]int{1}, []int{2}), "1,2"},
		{"", fn([]int{2}, []int{1}), "1,2"},
		{"", fn([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}), "1,1,2,2,3,3,4,4,5,5"},
		{"", fn([]int{1, 2, 3, 4}, []int{2, 4, 5, 6}), "1,2,2,3,4,4,5,6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (*LinkList)(MergeSortLink(tt.args.l1, tt.args.l2, cmp)).String(); got != tt.want {
				t.Errorf("MergeSortLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
