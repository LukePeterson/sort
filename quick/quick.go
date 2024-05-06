package quick

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
