package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/algo/sort/bs"
	"github.com/iamsumit/go-dsa/algo/sort/hs"
	"github.com/iamsumit/go-dsa/algo/sort/is"
	"github.com/iamsumit/go-dsa/algo/sort/ms"
	"github.com/iamsumit/go-dsa/algo/sort/qs"
	"github.com/iamsumit/go-dsa/algo/sort/ss"
)

func main() {
	fmt.Println("----------------------------")
	fmt.Println("------- Bubble Sort --------")
	fmt.Println("----------------------------")

	list := [][]int{
		{12, 11, 13, 5, 6},
		{4, 3, 2, 1, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for _, arr := range list {
		fmt.Println("Unsorted:", arr)
		bs.Sort(arr)
		fmt.Println("Sorted:", arr)
		fmt.Println("----------------------------")
	}

	fmt.Println("------ Selection Sort ------")
	fmt.Println("----------------------------")

	list = [][]int{
		{12, 11, 13, 5, 6},
		{4, 3, 2, 1, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for _, arr := range list {
		fmt.Println("Unsorted:", arr)
		ss.Sort(arr)
		fmt.Println("Sorted:", arr)
		fmt.Println("----------------------------")
	}

	fmt.Println("------ Insertion Sort ------")
	fmt.Println("----------------------------")
	list = [][]int{
		{12, 11, 13, 5, 6},
		{4, 3, 2, 1, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for _, arr := range list {
		fmt.Println("Unsorted:", arr)
		is.Sort(arr)
		fmt.Println("Sorted:", arr)
		fmt.Println("----------------------------")
	}

	fmt.Println("-------- Quick Sort --------")
	fmt.Println("----------------------------")

	list = [][]int{
		{12, 11, 13, 5, 6},
		{4, 3, 2, 1, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for _, arr := range list {
		fmt.Println("Unsorted:", arr)
		qs.Sort(arr)
		fmt.Println("Sorted:", arr)
		fmt.Println("----------------------------")
	}

	fmt.Println("-------- Merge Sort --------")
	fmt.Println("----------------------------")

	list = [][]int{
		{12, 11, 13, 5, 6},
		{4, 3, 2, 1, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for _, arr := range list {
		fmt.Println("Unsorted:", arr)
		ms.Sort(arr)
		fmt.Println("Sorted:", arr)
		fmt.Println("----------------------------")
	}

	fmt.Println("-------- Heap Sort ---------")
	fmt.Println("----------------------------")

	list = [][]int{
		{12, 11, 13, 5, 6, 7},
		{12, 5, 6, 7, 11, 13},
	}

	for _, arr := range list {
		fmt.Println("Unsorted:", arr)
		hs.Sort(arr)
		fmt.Println("Sorted:", arr)
		fmt.Println("----------------------------")
	}

}
