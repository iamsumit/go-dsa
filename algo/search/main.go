package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/algo/search/bs"
	"github.com/iamsumit/go-dsa/algo/search/is"
	"github.com/iamsumit/go-dsa/algo/search/js"
	"github.com/iamsumit/go-dsa/algo/search/ls"
)

func main() {
	fmt.Println("---------------------------------")
	fmt.Println("--------- Linear Search ---------")
	fmt.Println("---------------------------------")

	list := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, l := range list {
		fmt.Printf("Searching %d in %v\n", l.target, l.nums)
		fmt.Printf("Index: %d\n", ls.Search(l.nums, l.target))
		fmt.Println("---------------------------------")
	}

	fmt.Println("--------- Binary Search ---------")
	fmt.Println("---------------------------------")

	list = []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, l := range list {
		fmt.Printf("Searching %d in %v\n", l.target, l.nums)
		fmt.Printf("Index: %d\n", bs.Search(l.nums, l.target))
		fmt.Println("---------------------------------")
	}

	fmt.Println("--------- Ternary Search --------")
	fmt.Println("---------------------------------")

	list = []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, l := range list {
		fmt.Printf("Searching %d in %v\n", l.target, l.nums)
		fmt.Printf("Index: %d\n", bs.Search(l.nums, l.target))
		fmt.Println("---------------------------------")
	}

	fmt.Println("--------- Jump Search -----------")
	fmt.Println("---------------------------------")

	list = []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, -1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, l := range list {
		fmt.Printf("Searching %d in %v\n", l.target, l.nums)
		fmt.Printf("Index: %d\n", js.Search(l.nums, l.target))
		fmt.Println("---------------------------------")
	}

	fmt.Println("----- Interpolation Search ------")
	fmt.Println("---------------------------------")

	list = []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, -1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, l := range list {
		fmt.Printf("Searching %d in %v\n", l.target, l.nums)
		fmt.Printf("Index: %d\n", is.Search(l.nums, l.target))
		fmt.Println("---------------------------------")
	}

	fmt.Println("------ Exponential Search -------")
	fmt.Println("---------------------------------")

	list = []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, -1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, l := range list {
		fmt.Printf("Searching %d in %v\n", l.target, l.nums)
		fmt.Printf("Index: %d\n", is.Search(l.nums, l.target))
		fmt.Println("---------------------------------")
	}
}
