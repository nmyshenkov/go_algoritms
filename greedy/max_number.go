package greedy

import (
	"sort"
	"strconv"
)

func maxNumberSearch(numbers []int) string {
	var prep string

	for i := 0; i < len(numbers); i++ {
		var maxKey, maxVal int
		for k, val := range numbers {
			if val > maxVal {
				maxKey, maxVal = k, val
			}
		}
		numbers[maxKey] = -1
		prep = prep + strconv.Itoa(maxVal)
	}
	return prep
}

func maxNumberSort(numbers []int) string {
	var prep string
	sort.Ints(numbers)

	for i := len(numbers) - 1; i >= 0; i-- {
		prep = prep + strconv.Itoa(numbers[i])
	}
	return prep
}
