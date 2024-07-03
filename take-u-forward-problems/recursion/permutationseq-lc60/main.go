package main

import (
	"fmt"
	"slices"
	"strconv"
)

// https://takeuforward.org/data-structure/find-k-th-permutation-sequence/
// second approach
func getPermutation(n int, k int) string {
	var str string
	for i := 1; i <= n; i++ {
		str += strconv.Itoa(i)
	}

	perms := []string{}

	var nextPermFunc func(str string, index int)
	nextPermFunc = func(str string, index int) {
		if index == len(str) {
			perms = append(perms, str)
			return
		}

		for i := index; i < len(str); i++ {
			newStr := []byte(str)
			newStr[i], newStr[index] = newStr[index], newStr[i]
			nextPermFunc(string(newStr), index+1)
			newStr[index], newStr[i] = newStr[i], newStr[index]
		}

	}

	nextPermFunc(str, 0)

	slices.Sort(perms)

	return perms[k-1]
}

// func getPermutation(n int, k int) string {
// 	nums := make([]int, n)
// 	for i := 1; i <= n; i++ {
// 		nums[i-1] = i
// 	}

// 	var nextPermFunc func(num []int, count int)

// 	nextPermFunc = func(nums []int, count int) {
// 		if count == k {
// 			return
// 		}

// 		i := len(nums) - 2
// 		for i > 0 && nums[i] >= nums[i+1] {
// 			i--
// 		}

// 		j := len(nums) - 1
// 		for j > 0 && nums[j] <= nums[i] {
// 			j--
// 		}

// 		right := j
// 		left := i

// 		for left >= i && left < right {
// 			if nums[left] < nums[right] {
// 				nums[left], nums[right] = nums[right], nums[left]
// 				slices.Sort(nums[i+1:])
// 				break
// 			}

// 			left--
// 		}

// 		nextPermFunc(nums, count+1)
// 	}

// 	nextPermFunc(nums, 1)

// 	perm := ""
// 	for _, v := range nums {
// 		perm += strconv.Itoa(v)
// 	}

// 	return perm
// }

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
}
