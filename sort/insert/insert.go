package insert

func sort(in []int) []int {

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
