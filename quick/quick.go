package quick

// quickSort sorts an array using the quick sort algorithm
//
// Picks an element as a pivot and partitions the array around the pivot, placing smaller elements before it and larger elements after it, then sorts the partitions recursively.  Highly efficient with good average-case performance, not stable, and works well in practice despite O(n^2) worst-case scenario.
func quickSort(input []int) []int {
	return quickSortHelper(input, 0, len(input)-1)
}

func quickSortHelper(input []int, low int, high int) []int {
	if low < high {
		partitionIndex := partition(input, low, high)
		quickSortHelper(input, low, partitionIndex-1)
		quickSortHelper(input, partitionIndex+1, high)
	}
	return input
}

func partition(input []int, low int, high int) int {
	pivotIndex := high
	pivot := input[pivotIndex]
	high-- // Move 'high' off the pivot

	for {
		// Move 'low' forward as long as elements are less than the pivot and within bounds
		for low <= high && input[low] < pivot {
			low++
		}

		// Move 'high' backward as long as elements are greater than the pivot and within bounds
		for low <= high && input[high] > pivot {
			high--
		}

		if low >= high {
			break // If pointers cross, partioning is done
		} else {
			// Swap elements at 'low' and 'high'
			input[low], input[high] = input[high], input[low]
			low++
			high--
		}
	}

	// Swap the pivot back into its correct position
	input[low], input[pivotIndex] = input[pivotIndex], input[low]
	return low
}
