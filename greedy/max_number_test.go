package greedy

import "testing"

func Test_maxNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{[]int{3, 2, 5, 9, 1}}, "95321"},
		{"Test 2", args{[]int{3, 2, 5, 9, 1, 2, 4}}, "9543221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberSearch(tt.args.numbers); got != tt.want {
				t.Errorf("maxNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxNumberSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{[]int{3, 2, 5, 9, 1}}, "95321"},
		{"Test 2", args{[]int{3, 2, 5, 9, 1, 2, 4}}, "9543221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberSort(tt.args.numbers); got != tt.want {
				t.Errorf("maxNumberSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

var benchData = []int{3, 2, 5, 9, 1, 2, 4}

func Benchmark_maxNumberSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxNumberSearch(benchData)
	}
}

func Benchmark_maxNumberSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxNumberSort(benchData)
	}
}
