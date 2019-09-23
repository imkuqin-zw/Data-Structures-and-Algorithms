/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-15 01:02:29
 * @LastEditTime: 2019-09-15 13:48:00
 * @LastEditors: Please set LastEditors
 */

package main

import (
	"testing"
)

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"1+1"}, 2},
		{"", args{"1+(1+2)"}, 4},
		{"", args{"1+1+2^10"}, 1026},
		{"", args{"1+1+2^10*2"}, 2050},
		{"", args{"1+1+2^10*(2+1)"}, 3074},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
