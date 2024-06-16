package is

func Sort(arr []int) {
	// Start from the second element
	for i := 1; i < len(arr); i++ {
		// The key is the element we are going to insert
		key := arr[i]

		// j is the index of the element before the key
		j := i - 1

		// Move elements of arr[0..i-1], that are greater than key, to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Insert the key at its correct position
		arr[j+1] = key
	}
}
