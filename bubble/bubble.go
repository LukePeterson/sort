package bubble

// bubbleSort sorts an array of ints using the bubble sort algorithm
//
// Best Case: O(n) - only one pass is needed if the array is already sorted.  Our 'swapped' variable will tell us whether this is the case.
// Avg Case: O(n^2) - two nested loops are used to traverse the array.
// Worst Case: O(n^2) — when the array is sorted in reverse.
// Space Complexity: O(1) - in-place sorting algorithm so no new memory needed
// Stable.
func bubbleSort(input []int) []int {

	n := len(input)
	for i := 0; i < n-1; i++ {

		// This will tell us whether we can quit early
		swapped := false

		// We don't need to keep checking the end of the array as
		// we know it's already sorted, so only go up to n - i.
		for j := 0; j < n-i-1; j++ {

			// Swap the elements if j is bigger than j+1, in otherwords,
			// if we're out of order, swap the two elements.
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				swapped = true
			}
		}

		// If we did a full pass in the j loop without swapping, we must already be sorted, so we can exit early
		if !swapped {
			break
		}
	}

	return input
}
