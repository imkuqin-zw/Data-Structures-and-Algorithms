package main

import (
	"reflect"
	"testing"
)

func Test_index(t *testing.T) {
	type args struct {
		src string
		sub string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				src: "ABCDABCEAAAABASABCDABCADABCDABCEAABCDABCEAAABASABCDABCAABLAKABCDABABCDABCEAAADSFDABCADABCDABCEAAABCDABCEAAABASABCDABCADABCDABCEAAABLAKABLAKK",
				sub: "ABCDABCEAAABASABCDABCADABCDABCEAAABLAK",
			},
			96,
		},
		{
			"test2",
			args{
				src: "ABCDABCEAGFDSGFDGFGJHGJFDJ",
				sub: "ABCE",
			},
			4,
		},
		{
			"test3",
			args{
				src: "ABCDABCEAGFDSGFDGFGJHGJFDJ",
				sub: "A",
			},
			0,
		},

		{
			"test4",
			args{
				src: "ABCDABCEAGFDSGFDGFGJHGJFDJ",
				sub: "AB",
			},
			0,
		},
		{
			"test5",
			args{
				src: "ABCDABCEAGFDSGFDGFGUHGLFDJ",
				sub: "DJ",
			},
			24,
		},
		{
			"test6",
			args{
				src: "ABCDABCEAGFDSGFDGFGUHGLFDJ",
				sub: "J",
			},
			25,
		},
		{
			"test7",
			args{
				src: "ABCDABCEAGFDSGFDGFGUHGLFDJ",
				sub: "",
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := index(tt.args.src, tt.args.sub); got != tt.want {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	type args struct {
		sub string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				sub: "ABCDABCEAAABASABCDABCADABCDABCEAAABLAK",
			},
			[]int{-1, 0, 0, 0, 0, 1, 2, 3, 0, 1, 1, 1, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.sub); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%d, %d\ngetNext()   %v \nwant \t\t%v", len(tt.args.sub), len(tt.want), got, tt.want)
			}
		})
	}
}

func Test_commStr(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"tset1",
			args{"abcdefgtgfderg", "cdefggfdsg"},
			"cdefggfdg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commStr(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("commStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
