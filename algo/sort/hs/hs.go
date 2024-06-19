package hs

func Sort(arr []int) []int {
	length := len(arr)

	for i := length/2 - 1; i >= 0; i-- {
		heapify(arr, length, i)
	}

	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}

func heapify(arr []int, length, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && arr[left] > arr[largest] {
		largest = left
	}

	if right < length && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, length, largest)
	}
}
