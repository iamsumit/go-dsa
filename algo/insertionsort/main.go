package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/algo/insertionsort/is"
)

func main() {
	list := [][]int{
		{12, 11, 13, 5, 6},
		{4, 3, 2, 1, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for _, arr := range list {
		fmt.Println("Unsorted:", arr)
		is.Sort(arr)
		fmt.Println("Sorted:", arr)
	}
}
