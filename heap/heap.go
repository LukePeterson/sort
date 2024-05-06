package heap

// heapSort sorts an array using the heap sort algorithm
func heapSort(input []int) []int {
	n := len(input)

	// Since last parent will be at n/2 - 1 we can start from here
	for i := n/2 - 1; i >= 0; i-- {
		heapify(input, n, i)
	}

	// One by one extract elements
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		input[0], input[i] = input[i], input[0]

		// Call max heapify on the reduced heap
		heapify(input, i, 0)
	}

	return input
}

// heapify function transforms the slice into a heap at a specified index in a slice of given length.
func heapify(input []int, n int, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // left = 2*i + 1
	right := 2*i + 2 // right = 2*i + 2

	// If left child is larger than root
	if left < n && input[left] > input[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && input[right] > input[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		input[i], input[largest] = input[largest], input[i] // Swap

		// Recursively heapify the affected sub-tree
		heapify(input, n, largest)
	}
}
