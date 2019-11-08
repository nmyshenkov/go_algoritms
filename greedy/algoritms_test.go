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
			if got := maxNumber(tt.args.numbers); got != tt.want {
				t.Errorf("maxNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
