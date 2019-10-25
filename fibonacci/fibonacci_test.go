package fibonacci

import "testing"

func Test_get(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "Test 1. First elem", args: args{n: 1}, want: 1},
		{name: "Test 2. Second elem", args: args{n: 2}, want: 1},
		{name: "Test 3. Fifth elem", args: args{n: 5}, want: 5},
		{name: "Test 4. Tenth elem", args: args{n: 10}, want: 55},
		{name: "Test 5. Twentieth elem", args: args{n: 20}, want: 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get(tt.args.n); got != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}
