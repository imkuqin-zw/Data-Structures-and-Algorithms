package heap

import (
	"fmt"
	"testing"
)

func cmp(a, b interface{}) int {
	if a.(int) == b.(int) {
		return 0
	} else if a.(int) > b.(int) {
		return 1
	} else {
		return -1
	}
}

func TestHeap_heap(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "", args: []int{1, 2, 3, 5, 9, 7}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(cmp, 0)
			for _, e := range tt.args {
				h.Insert(e)
			}
			fmt.Println(h.elems[1:])
			fmt.Print("delete: ")
			for {
				result := h.DelMax()
				if result == nil {
					break
				}
				fmt.Print(result, " ")
			}
			fmt.Println()
			fmt.Println(h.elems[1:])
		})
	}
}

func TestHeap_Sort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "", args: []int{}, want: []int{}},
		{name: "", args: []int{1, 2, 3, 5, 9, 7}, want: []int{1, 2, 3, 5, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(cmp, 0)
			for _, e := range tt.args {
				h.Insert(e)
			}
			got := h.Sort()
			for i, num := range got {
				if num != tt.want[i] {
					t.Errorf("Sort() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
