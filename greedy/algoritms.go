package greedy

import "strconv"

func maxNumber(numbers []int) string {
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
