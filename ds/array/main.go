package main

import "fmt"

func main() {
	fmt.Println("--------------------")
	numSeries := [5]int{1, 2, 3, 4, 5}

	fmt.Println("All elements:", numSeries)
	fmt.Println("First element:", numSeries[0])
	fmt.Println("Length:", len(numSeries))

	fmt.Println("--------------------")

	numSeries[0] = 10
	fmt.Println("First element after updating its value to 10:", numSeries[0])
	fmt.Println("All elements now:", numSeries)

	fmt.Println("--------------------")

	numSeries = [5]int{10, 20, 30, 40, 50}
	fmt.Println("All elements after updating the array:", numSeries)
	fmt.Println("Last element:", numSeries[len(numSeries)-1])
	fmt.Println("Second and third element:", numSeries[1:3])
	fmt.Println("All elements after 2nd element:", numSeries[1:])
	fmt.Println("All elements before 2nd element:", numSeries[:1])

	fmt.Println("--------------------")

	fmt.Println("Iterating over the array:")
	for i, num := range numSeries {
		fmt.Println("Index:", i, "Value:", num)
	}

	fmt.Println("--------------------")

	element := 30
	fmt.Println("Element", element, "found in the array:", findElementInArray(numSeries, element))

	element = 100
	fmt.Println("Element", element, "found in the array:", findElementInArray(numSeries, element))

	fmt.Println("--------------------")

	matrix := [2][2]int{{11, 12}, {21, 22}}
	fmt.Println("Matrix:", matrix)
	fmt.Println("First row:", matrix[0])
	fmt.Println("Second row:", matrix[1])

	fmt.Println("First element of the first row:", matrix[0][0])
	fmt.Println("Second element of the first row:", matrix[0][1])

	fmt.Println("--------------------")

	matrix[0] = [2]int{110, 120}
	fmt.Println("Matrix after updating the first row:", matrix)

	fmt.Println("--------------------")
}

func findElementInArray(numSeries [5]int, element int) bool {
	for _, num := range numSeries {
		if num == element {
			return true
		}
	}
	return false
}
