package binary_tree

import (
	"reflect"
	"testing"
)

func TestBNode_LevelTraversal(t *testing.T) {
	type fields struct {
		node *BNode
	}
	fn := func(data []interface{}) fields {
		return fields{(*BNode)(CreateBTree(data, 0, len(data)-1))}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{"", fn([]interface{}{}), []interface{}{}},
		{"", fn([]interface{}{1}), []interface{}{1}},
		{"", fn([]interface{}{1, 2}), []interface{}{2, 1}},
		{"", fn([]interface{}{1, 2, 3}), []interface{}{2, 1, 3}},
		{"", fn([]interface{}{1, 2, 3, 4, 5, 6, 7}), []interface{}{4, 2, 6, 1, 3, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.fields.node
			if got := b.LevelTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBNode_PreOrderTraversal(t *testing.T) {
	type fields struct {
		node *BNode
	}
	fn := func(data []interface{}) fields {
		return fields{(*BNode)(CreateBTree(data, 0, len(data)-1))}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{"", fn([]interface{}{}), []interface{}{}},
		{"", fn([]interface{}{1}), []interface{}{1}},
		{"", fn([]interface{}{1, 2}), []interface{}{2, 1}},
		{"", fn([]interface{}{1, 2, 3}), []interface{}{2, 1, 3}},
		{"", fn([]interface{}{1, 2, 3, 4, 5, 6, 7}), []interface{}{4, 2, 1, 3, 6, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.fields.node
			if got := b.PreOrderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBNode_InOrderTraversal(t *testing.T) {
	type fields struct {
		node *BNode
	}
	fn := func(data []interface{}) fields {
		return fields{(*BNode)(CreateBTree(data, 0, len(data)-1))}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{"", fn([]interface{}{}), []interface{}{}},
		{"", fn([]interface{}{1}), []interface{}{1}},
		{"", fn([]interface{}{1, 2}), []interface{}{1, 2}},
		{"", fn([]interface{}{1, 2, 3}), []interface{}{1, 2, 3}},
		{"", fn([]interface{}{1, 2, 3, 4, 5, 6, 7}), []interface{}{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.fields.node
			if got := b.InOrderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBNode_PostOrderTraversal(t *testing.T) {
	type fields struct {
		node *BNode
	}
	fn := func(data []interface{}) fields {
		return fields{(*BNode)(CreateBTree(data, 0, len(data)-1))}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{"", fn([]interface{}{}), []interface{}{}},
		{"", fn([]interface{}{1}), []interface{}{1}},
		{"", fn([]interface{}{1, 2}), []interface{}{1, 2}},
		{"", fn([]interface{}{1, 2, 3}), []interface{}{1, 3, 2}},
		{"", fn([]interface{}{1, 2, 3, 4, 5, 6, 7}), []interface{}{1, 3, 2, 5, 7, 6, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.fields.node
			if got := b.PostOrderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
