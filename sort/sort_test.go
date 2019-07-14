package sort

import (
	"reflect"
	"testing"
)

type args struct {
	in []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{name: "test 1", args: args{in: []int{1, 4, 2}}, want: []int{1, 2, 4}},
	{name: "test 2", args: args{in: []int{2, 8, 4, 1}}, want: []int{1, 2, 4, 8}},
	{name: "test 3", args: args{in: []int{3, 7, 4, 1, 4, 8}}, want: []int{1, 3, 4, 4, 7, 8}},
}

func Test_bubble(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubble(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

var benchData = []int{1, 31, 4, 6, 7, 4, 10, 12, 9}

func Benchmark_bubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble(benchData)
	}
}

func Benchmark_sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insert(benchData)
	}
}
