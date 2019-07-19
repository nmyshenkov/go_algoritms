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

func merge(in []int) []int {

	num := len(in)

	if num == 1 {
		return in
	}

	middle := int(num / 2)

	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)

	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = in[i]
		} else {
			right[i-middle] = in[i]
		}
	}

	return m(merge(left), merge(right))
}

func m(left, right []int) (r []int) {
	r = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			r[i] = left[0]
			left = left[1:]
		} else {
			r[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		r[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		r[i] = right[j]
		i++
	}

	return r
}
