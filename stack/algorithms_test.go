package stack

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/arrayStack"
	"reflect"
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

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evalRPN(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{"", []string{"2", "1", "+", "3", "*"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslands(t *testing.T) {
	tests := []struct {
		name string
		args [][]byte
		want int
	}{
		{"", [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{nums: []int{1, 1, 1, 1, 1}, S: 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dominantIndex(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		//{"", []int{0,0,0,1}, 3},
		{"", []int{0, 0, 1, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.args); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
