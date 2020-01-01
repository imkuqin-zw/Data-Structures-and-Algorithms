package listQueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name string
		elms []interface{}
		want [][]interface{}
	}{
		{
			name: "",
			elms: []interface{}{1, 2, 3, 4, 5, 6},
			want: [][]interface{}{
				{1, 2, 3, 4, 5, 6},
				{2, 3, 4, 5, 6},
				{3, 4, 5, 6},
				{4, 5, 6},
				{5, 6},
				{6},
				{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, item := range tt.elms {
				q.Enqueue(item)
			}
			i := 0
			for !q.Empty() {
				assert.Equal(t, tt.want[i], q.traversal())
				i++
				q.Dequeue()
			}
			assert.Equal(t, tt.want[i], q.traversal())
		})
	}
}
