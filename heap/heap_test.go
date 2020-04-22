package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestHeap_PushPop(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []interface{}
	}{
		{name: "", args: []int{1, 2, 3, 5, 9, 7}, want: []interface{}{9, 7, 5, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(cmp, 0)
			for _, e := range tt.args {
				h.Push(e)
			}
			result := make([]interface{}, 0)
			for {
				e := h.Pop()
				if e == nil {
					break
				}
				result = append(result, e)
			}
			assert.Equal(t, result, tt.want)
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
				h.Push(e)
			}
			got := h.Sort()
			assert.Equal(t, tt.want, got)
		})
	}
}
