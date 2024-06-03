package main

import "fmt"

func main() {
	fmt.Println("--------------------")

	numSeries := make([]int, 2, 5)
	fmt.Println("All elements:", numSeries)
	// Length: 2
	fmt.Println("Length:", len(numSeries))
	// Capacity: 5
	fmt.Println("Capacity:", cap(numSeries))

	fmt.Println("--------------------")

	numSeries = []int{1, 2, 3, 4, 5}

	fmt.Println("All elements:", numSeries)
	fmt.Println("First element:", numSeries[0])
	// Length: 5
	fmt.Println("Length:", len(numSeries))
	// Capacity: 5
	fmt.Println("Capacity:", cap(numSeries))

	fmt.Println("--------------------")

	numSeries = append(numSeries, 6)
	numSeries = append(numSeries, 7)

	fmt.Println("All elements:", numSeries)
	// Length: 7
	fmt.Println("Length:", len(numSeries))
	// Capacity: 10 - doubles the capacity when the slice is full
	fmt.Println("Capacity:", cap(numSeries))

	fmt.Println("--------------------")

	numSeries = append(numSeries, 8)
	numSeries = append(numSeries, 9)
	numSeries = append(numSeries, 10)

	fmt.Println("All elements:", numSeries)
	// Length: 10
	fmt.Println("Length:", len(numSeries))
	// Capacity: 10
	fmt.Println("Capacity:", cap(numSeries))

	fmt.Println("--------------------")

	numSeries = append(numSeries, 11)
	fmt.Println("All elements:", numSeries)
	// Length: 11
	fmt.Println("Length:", len(numSeries))
	// Capacity: 20 - doubles the capacity when the slice is full
	fmt.Println("Capacity:", cap(numSeries))

	fmt.Println("--------------------")

	fmt.Println("Iterating over the slice:")
	for i, num := range numSeries {
		fmt.Println("Index:", i, "Value:", num)
	}

	fmt.Println("--------------------")

	element := 1
	fmt.Println("Element", element, "found in the slice:", findElementInSlice(numSeries, element))

	element = 12
	fmt.Println("Element", element, "found in the slice:", findElementInSlice(numSeries, element))

	fmt.Println("--------------------")

	numSeries = append(numSeries[:2], numSeries[4:]...)

	fmt.Println("All elements after removing the third and fourth element:", numSeries)
	// Length: 9
	fmt.Println("Length:", len(numSeries))
	// Capacity: 20 - it still maintains the last capacity.
	fmt.Println("Capacity:", cap(numSeries))

	fmt.Println("--------------------")
}

func findElementInSlice(numSeries []int, element int) bool {
	for _, num := range numSeries {
		if num == element {
			return true
		}
	}
	return false
}
