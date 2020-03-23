package listStack

import (
	"fmt"
	"testing"
)

//func TestStack_Top(t *testing.T) {
//	type fields struct {
//		top    *Node
//		length int
//	}
//	node2 := NewNode(2)
//	node2.Next = NewNode(1)
//	tests := []struct {
//		name   string
//		fields fields
//		want   interface{}
//	}{
//		{"", fields{length: 0}, nil},
//		{"", fields{top: NewNode(1), length: 1}, 1},
//		{"", fields{top: node2, length: 2}, 2},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Stack{
//				top:    tt.fields.top,
//				length: tt.fields.length,
//			}
//			assert.Equal(t, tt.want, s.Top())
//		})
//	}
//}

//
//func TestStack_Pop(t *testing.T) {
//	type fields struct {
//		top    *Node
//		length int
//	}
//	type want struct {
//		top    interface{}
//		length int
//	}
//	node2 := NewNode(2)
//	node2.Next = NewNode(1)
//	tests := []struct {
//		name   string
//		fields fields
//		want   want
//	}{
//		{"", fields{length: 0}, want{nil, 0}},
//		{"", fields{top: NewNode(1), length: 1}, want{nil, 0}},
//		{"", fields{top: node2, length: 2}, want{1, 1}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Stack{
//				top:    tt.fields.top,
//				length: tt.fields.length,
//			}
//			s.Pop()
//			if assert.Equal(t, tt.want.length, s.Length()) {
//				assert.Equal(t, tt.want.top, s.Top())
//			}
//		})
//	}
//}

func TestStack_Push(t *testing.T) {
	fn := func(vals ...interface{}) {
		s := &Stack{}
		for _, val := range vals {
			s.Push(val)
		}
		s.Traversal()
	}
	fn(1, 2, 3, 4, 5, 6)
	fn(7, 8, 9, 10, 11, 12, 13)
}

func TestStack_Pop(t *testing.T) {
	fn := func(vals ...interface{}) {
		s := &Stack{}
		for _, val := range vals {
			s.Push(val)
		}
		s.Traversal()
		for {
			data := s.Pop()
			if data == nil {
				break
			}
			fmt.Print(data, " ")
		}
		fmt.Println()
		fmt.Println(s.length)
		s.Traversal()
	}
	fn(1, 2, 3, 4, 5, 6)
	//fn(7, 8, 9, 10, 11, 12, 13)
}
