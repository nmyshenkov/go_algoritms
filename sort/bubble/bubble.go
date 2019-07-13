package bubble

func sort(in []int) []int {

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
