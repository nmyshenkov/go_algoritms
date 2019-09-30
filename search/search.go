package search

// linear - linear search in int array
func linear(el int, arr []int) bool {

	for _, val := range arr {
		if el == val {
			return true
		}
	}
	return false
}

// binary - binary search in sorted int array
func binary(el int, arr []int) bool {
	// check if arr len eq 1
	if len(arr) == 1 {
		if arr[0] == el {
			return true
		}
		return false
	}

	var (
		left, mid, right int
	)

	right = len(arr)

	// check sort type and flip if it need
	if arr[0] > arr[len(arr)-1] {
		for i := len(arr)/2 - 1; i >= 0; i-- {
			opp := len(arr) - 1 - i
			arr[i], arr[opp] = arr[opp], arr[i]
		}
	}

	for !(left >= right) {
		mid = left + (right-left)/2

		if arr[mid] == el {
			return true
		}

		if arr[mid] > el {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return false
}
