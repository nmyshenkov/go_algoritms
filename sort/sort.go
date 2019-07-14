package sort

func bubble(in []int) []int {

	n := len(in)

	for j := 0; j < n-1; j++ {

		for i := 0; i < n-j-1; i++ {
			if in[i] > in[i+1] {
				in[i], in[i+1] = in[i+1], in[i]
			}
		}
	}

	return in
}

func insert(in []int) []int {

	for i := 1; i < len(in); i++ {

		value := in[i]
		o := i

		for j := i - 1; j > -1; j-- {
			if in[j] < value {
				break
			}
			in[j+1] = in[j]
			o = j
		}
		in[o] = value
	}

	return in
}
