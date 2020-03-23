package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{data: []int{}}, []int{}},
		{"", args{data: []int{1}}, []int{1}},
		{"", args{data: []int{9, 6, 5, 7, 8, 1, 3, 2, 4}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectSort(tt.args.data)
			assert.Equal(t, tt.want, tt.args.data)
		})
	}
}

func TestInsertSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{data: []int{}}, []int{}},
		{"", args{data: []int{1}}, []int{1}},
		{"", args{data: []int{9, 6, 5, 7, 8, 1, 3, 2, 4}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.data)
			assert.Equal(t, tt.want, tt.args.data)
		})
	}
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{data: []int{}}, []int{}},
		{"", args{data: []int{1}}, []int{1}},
		{"", args{data: []int{9, 6, 5, 7, 8, 1, 3, 2, 4}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.data)
			assert.Equal(t, tt.want, tt.args.data)
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{data: []int{}}, []int{}},
		{"", args{data: []int{1}}, []int{1}},
		{"", args{data: []int{9, 6, 5, 7, 8, 1, 3, 2, 4}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.data, 0, len(tt.args.data)-1)
			assert.Equal(t, tt.want, tt.args.data)
		})
	}
}
