package selection

// selectionSort sorts an array of ints using the selection sort algorithm
//
// Finds the smallest (or largest) element in the array and moves it to the start, then repeats for the remainder of the list.  Unstable and performs poorly compared to other sorts, but easy to implement.
//
// Best Case: O(n^2) — regardless of initial order of elements.
// Average Case: O(n^2)
// Worst Case: O(n^2)
// Space Complexity: O(1) - in-place sorting algorithm so no new memory needed
// NOT Stable.
func selectionSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		minIndex := i // Start our minimum as whatever is at the first element
		for j := i + 1; j < len(input); j++ {
			// If our current j element is less than what we've found on this pass of j
			if input[j] < input[minIndex] {
				minIndex = j // Set the new minimum index
			}
		}

		// We found a value that's lower than what's at index[i], so ...
		if minIndex != i {
			input[i], input[minIndex] = input[minIndex], input[i] // ... swap them
		}
	}

	return input
}
