package merge

// mergeSort is the main function that implements the merge sort algorithm.
// It recursively splits the array into halves, sorts each half, and merges them back together.
func mergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	// Find the middle point to divide the array into two halves
	mid := len(input) / 2

	// Recursively sort the first half
	left := mergeSort(input[:mid])

	// Recursively sort the second half
	right := mergeSort(input[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// merge is a helper function that merges two halves of an array.
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	// Continue looping until we reach the end of either `left` or `right`
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append the remaining elements of `left` (if any)
	for ; i < len(left); i++ {
		result = append(result, left[i])
	}

	// Append the remaining elements of `right` (if any)
	for ; j < len(right); j++ {
		result = append(result, right[j])
	}

	return result
}
