package main

import "fmt"

func main() {
	fmt.Println("---------------------------------")
	fmt.Println("Implementing a map in Go")
	fmt.Println("---------------------------------")

	mapVar := map[int]int{}
	mapVar[1] = 1

	fmt.Println("mapVar[1]:", mapVar[1])

	mapVar[2] = 2
	fmt.Println("mapVar[2]:", mapVar[2])

	fmt.Println("len(mapVar):", len(mapVar))

	fmt.Println("---------------------------------")

	delete(mapVar, 1)
	fmt.Println("mapVar[1] after deleting:", mapVar[1])

	_, ok := mapVar[1]
	fmt.Println("mapVar[1] exists:", ok)

	_, ok = mapVar[2]
	fmt.Println("mapVar[2] exists:", ok)

	fmt.Println("len(mapVar):", len(mapVar))

	fmt.Println("---------------------------------")

	mapVar[2] = 3
	fmt.Println("mapVar[2] after updating:", mapVar[2])

	fmt.Println("---------------------------------")

	mapVar[1] = 1
	mapVar[2] = 1
	mapVar[3] = 3
	fmt.Println("Iterating over mapVar:")
	for key, value := range mapVar {
		fmt.Println("key:", key, "value:", value)
	}

	fmt.Println("---------------------------------")
}
