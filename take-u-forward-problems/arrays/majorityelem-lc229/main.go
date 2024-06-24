package main

import (
	"fmt"
)

// Moore's Voting Algo
func majorityElement(nums []int) []int {
	var count1, count2, elem1, elem2 int
	for _, n := range nums {
		if count1 == 0 && n != elem2 {
			count1 = 1
			elem1 = n
		} else if count2 == 0 && n != elem1 {
			count2 = 1
			elem2 = n
		} else if elem1 == n {
			count1++
		} else if elem2 == n {
			count2++
		} else {
			count1--
			count2--
		}
	}

	if elem1 == elem2 {
		return []int{elem1}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == elem1 {
			count1++
		}

		if num == elem2 {
			count2++
		}
	}

	elems := []int{}
	if count1 > len(nums)/3 {
		elems = append(elems, elem1)
	}

	if count2 > len(nums)/3 {
		elems = append(elems, elem2)
	}

	return elems
}

// With hashmap
// func majorityElement(nums []int) []int {
// 	hashMap := map[int]int{}

// 	elems := []int{}

// 	for _, num := range nums {
// 		hashMap[num]++
// 		if hashMap[num] == len(nums)/3+1 {
// 			elems = append(elems, num)
// 		}

// 		if len(elems) == 2 {
// 			break
// 		}

// 	}

// 	return elems
// }

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 3, 2, 2}))
	fmt.Println(majorityElement([]int{2, 2, 1, 3}))
}
