package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Top(t *testing.T) {
	type fields struct {
		top    *Node
		length int
	}
	node2 := NewNode(2)
	node2.Next = NewNode(1)
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"", fields{length: 0}, nil},
		{"", fields{top: NewNode(1), length: 1}, 1},
		{"", fields{top: node2, length: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			assert.Equal(t, tt.want, s.Top())
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		top    *Node
		length int
	}
	type args struct {
		val interface{}
	}
	node2 := NewNode(2)
	node2.Next = NewNode(1)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{length: 0}, args{1}},
		{"", fields{top: NewNode(1), length: 1}, args{2}},
		{"", fields{top: node2, length: 2}, args{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			s.Push(tt.args.val)
			if assert.Equal(t, tt.fields.length+1, s.Length()) {
				assert.Equal(t, tt.args.val, s.Top())
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		top    *Node
		length int
	}
	type want struct {
		top    interface{}
		length int
	}
	node2 := NewNode(2)
	node2.Next = NewNode(1)
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{"", fields{length: 0}, want{nil, 0}},
		{"", fields{top: NewNode(1), length: 1}, want{nil, 0}},
		{"", fields{top: node2, length: 2}, want{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			s.Pop()
			if assert.Equal(t, tt.want.length, s.Length()) {
				assert.Equal(t, tt.want.top, s.Top())
			}
		})
	}
}
