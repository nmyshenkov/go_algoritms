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

func choose(in []int) []int {

	var (
		min int
	)

	for i := 0; i < len(in); i++ {
		min = i

		for j := i + 1; j < len(in); j++ {
			if in[j] < in[min] {
				min = j
			}
		}

		in[min], in[i] = in[i], in[min]
	}

	return in
}

func quick(in []int) []int {

	if len(in) < 2 {
		return in
	}

	l, r := 0, len(in)-1

	o := (l + r) / 2

	in[o], in[r] = in[r], in[o]

	for i := range in {
		if in[i] < in[r] {
			in[l], in[i] = in[i], in[l]
			l++
		}
	}

	in[l], in[r] = in[r], in[l]

	quick(in[:l])
	quick(in[l+1:])

	return in
}
