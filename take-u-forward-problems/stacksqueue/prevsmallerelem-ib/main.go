package main

import "fmt"

// https://www.interviewbit.com/problems/nearest-smaller-element/
func prevSmaller(nums []int) []int {
	stack := []int{}
	set := map[int]int{}

	for i, num := range nums {
		for len(stack) > 0 {
			if stack[len(stack)-1] < num {
				set[i] = stack[len(stack)-1]
				break
			}

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num)
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		val, ok := set[i]
		if !ok {
			result = append(result, -1)
		} else {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	fmt.Println(prevSmaller([]int{4, 5, 2, 10, 8}))
	fmt.Println(prevSmaller([]int{3, 2, 1}))
}
