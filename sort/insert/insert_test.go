package insert

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test 1", args: args{in: []int{1, 4, 2}}, want: []int{1, 2, 4}},
		{name: "test 2", args: args{in: []int{2, 8, 4, 1}}, want: []int{1, 2, 4, 8}},
		{name: "test 3", args: args{in: []int{3, 7, 4, 1, 4, 8}}, want: []int{1, 3, 4, 4, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
