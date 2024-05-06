package insertion

// insertionSort sorts an array of ints using the insertion sort algorithm
//
// Best Case: O(n) — when the array is already sorted.
// Average Case: O(n^2) — each element may need to be compared with all others already sorted.
// Worst Case: O(n^2) — when the array is sorted in reverse.
// Space Complexity: O(1) — an in-place sorting algorithm (like bubble sort)
// Stable.
func insertionSort(input []int) []int {

	for i := 1; i < len(input); i++ {
		currentValue := input[i]
		pos := i - 1 // Store the correct position for this element in the array

		// Shift elements of the sorted segment upwards to make space
		for pos >= 0 {
			if input[pos] > currentValue {
				input[pos+1] = input[pos] // Create an empty spot at input[pos]
				pos--
			} else {
				break
			}
		}
		input[pos+1] = currentValue // Place the current element in the correct position
	}

	return input
}
