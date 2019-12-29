package linklisk

import (
	"testing"
)

func TestReverse(t *testing.T) {
	h := new(ListNode)
	for i := 0; i < 10; i++ {
		HeadInsert(h, i)
	}
	Traverse(h)
	Reverse(h)
	Traverse(h)
}

func TestRemoveDup(t *testing.T) {
	h := new(ListNode)
	HeadInsert(h, 1)
	HeadInsert(h, 2)
	HeadInsert(h, 1)
	HeadInsert(h, 3)
	HeadInsert(h, 4)
	HeadInsert(h, 2)
	HeadInsert(h, 5)
	Traverse(h)
	RemoveDup(h)
	Traverse(h)
}

func TestAdd(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	f := func(vals [][]int) args {
		lists := make([]*ListNode, 2)
		for i, val := range vals {
			lists[i] = new(ListNode)
			for _, item := range val {
				TailInsert(lists[i], item)
			}
		}
		return args{l1: lists[0], l2: lists[1]}
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: f([][]int{{3, 1, 5}, {5, 9, 2}}), want: "808"},
		{name: "2", args: f([][]int{{1, 2}, {3, 2, 1}}), want: "441"},
		{name: "3", args: f([][]int{{}, {}}), want: "0"},
		{name: "4", args: f([][]int{{1}, {}}), want: "1"},
		{name: "5", args: f([][]int{{}, {1}}), want: "1"},
		{name: "6", args: f([][]int{{1}, {1}}), want: "2"},
		{name: "7", args: f([][]int{{3, 4, 5, 6, 7, 8}, {9, 8, 7, 6, 5}}), want: "233339"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(Add(tt.args.l1, tt.args.l2)); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
