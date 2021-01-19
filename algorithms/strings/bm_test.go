package strings

import "testing"

func TestBmIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{s: "aaaaa", substr: "bba"}, -1},
		{"", args{s: "mississippi", substr: "pi"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BmIndex(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("BmIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
