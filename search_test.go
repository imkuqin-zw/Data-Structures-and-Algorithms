package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{0}, 0},
		{"", args{1}, 1},
		{"", args{2}, 1},
		{"", args{3}, 2},
		{"", args{4}, 3},
		{"", args{5}, 5},
		{"", args{6}, 8},
		{"", args{7}, 13},
		{"", args{8}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
