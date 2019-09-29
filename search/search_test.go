package search

import "testing"

var arr = []int{3, 15, 23, 51, 20, 10, 11, 65, 5, 8, 0, 22}

func Test_linear(t *testing.T) {
	type args struct {
		el  int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test 1", args: args{el: 8, arr: arr}, want: true},
		{name: "Test 2", args: args{el: 7, arr: arr}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linear(tt.args.el, tt.args.arr); got != tt.want {
				t.Errorf("linear() = %v, want %v", got, tt.want)
			}
		})
	}
}
