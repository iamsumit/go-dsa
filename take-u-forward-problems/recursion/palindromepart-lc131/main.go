package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/palindrome-partitioning/
func partition(s string) [][]string {
	subsets := [][]string{}
	sets := []string{}

	var recurFunc func(index int)
	recurFunc = func(index int) {
		if index == len(s) {
			subsets = append(subsets, append([]string{}, sets...))
			return
		}

		for i := index; i < len(s); i++ {
			if isPalindrome(s[index : i+1]) {
				sets = append(sets, s[index:i+1])
				recurFunc(i + 1)
				sets = sets[:len(sets)-1]
			}
		}
	}

	recurFunc(0)

	return subsets
}

func isPalindrome(sets string) bool {
	left, right := 0, len(sets)-1
	for left <= right {
		if sets[left] != sets[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
	fmt.Println(partition("aabb"))
}
