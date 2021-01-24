package tree

import "testing"

func TestBaiLian4080(t *testing.T) {
	type args struct {
		val []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{val: []int{1, 1, 3, 5}}, 17},
		{"", args{val: []int{}}, 0},
		{"", args{val: []int{1}}, 1},
		{"", args{val: []int{1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaiLian4080(tt.args.val); got != tt.want {
				t.Errorf("BaiLian4080() = %v, want %v", got, tt.want)
			}
		})
	}
}
