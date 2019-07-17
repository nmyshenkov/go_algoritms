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
	// {name: "test 4", args: args{
	// 	in: []int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}},
	// 	want: []int{-590, -482, -412, -392, -381, -317, -312, -243, -215, -46, -14, 22, 158, 177, 217, 273, 380, 424, 514, 925}},
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

func Test_choose(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := choose(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("choose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quick(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quick(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quick() = %v, want %v", got, tt.want)
			}
		})
	}
}

//var benchData = []int{1, 31, 4, 6, 7, 4, 10, 12, 9}
var benchData = []int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}

func Benchmark_bubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble(benchData)
	}
}

func Benchmark_insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insert(benchData)
	}
}

func Benchmark_choose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		choose(benchData)
	}
}

func Benchmark_quick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quick(benchData)
	}
}
