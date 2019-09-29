package search

// linear - linear search on int array
func linear(el int, arr []int) bool {

	for _, val := range arr {
		if el == val {
			return true
		}
	}
	return false
}
