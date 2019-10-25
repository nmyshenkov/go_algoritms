package fibonacci

func get(n int) int64 {
	// make array with fix capacity
	series := make([]int64, n+1)
	series[0] = 0
	series[1] = 1

	for i := 2; i <= n; i++ {
		series[i] = series[i-1] + series[i-2]
	}

	return series[n]
}
