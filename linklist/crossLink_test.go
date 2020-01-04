package linklist

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	type args struct {
		l *CrossNode
	}
	cmp := func(val1, val2 interface{}) int {
		return val1.(int) - val2.(int)
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{
			CreateList([][]interface{}{{1, 1, 2, 3}}),
		}, "1,1,2,3"},
		{"", args{
			CreateList([][]interface{}{
				{1, 1, 2, 3},
				{1, 2, 4, 5, 8, 9},
				{2, 3, 5, 6, 8},
			}),
		}, "1,1,1,2,2,2,3,3,4,5,5,6,8,8,9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.l, cmp).String(); got != tt.want {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
