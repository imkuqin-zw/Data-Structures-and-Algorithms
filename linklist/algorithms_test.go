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
		{"", fn(1, 2, 3, 4, 5, 6, 7, 8), "1,8,2,7,3,6,4,5"},
		{"", fn(1, 2, 3, 4, 5, 6, 7), "1,7,2,6,3,5,4"},
		{"", fn(), ""},
		{"", fn(1), "1"},
		{"", fn(1, 2), "1,2"},
		{"", fn(1, 2, 3), "1,3,2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (*LinkList)(Sort1(tt.args.l1)).String(); got != tt.want {
				t.Errorf("Sort1() = %v, want %v", got, tt.want)
			}
		})
	}
}
