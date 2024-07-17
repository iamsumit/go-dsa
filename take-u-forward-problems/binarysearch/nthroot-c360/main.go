package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/nth-root-of-a-number-using-binary-search/

// approach 3
func nthRoot(n int, m int) int {
	low := 1
	high := m

	calc := func(mid, n, m int) int {
		ans := 1

		for i := 1; i <= n; i++ {
			ans *= mid
			if ans > m {
				return 2
			}
		}

		if ans == m {
			return 1
		}

		return 0
	}

	for low <= high {
		mid := (low + high) / 2
		midN := calc(mid, n, m)
		if midN == 1 {
			return mid
		} else if midN == 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// Brute force 2
// func nthRoot(n int, m int) int {
// 	calc := func(b, exp int) int {
// 		ans := 1
// 		base := b

// 		for exp > 0 {
// 			if exp%2 != 0 {
// 				exp -= 1
// 				ans = ans * base
// 			} else {
// 				exp /= 2
// 				base = base * base
// 			}
// 		}

// 		return ans
// 	}

// 	i := 1
// 	for {
// 		res := calc(i, n)
// 		if res == m {
// 			return i
// 		} else if res > m {
// 			break
// 		}

// 		i++
// 	}

// 	return -1
// }

// Brute force approach
// func nthRoot(n int, m int) int {
// 	if m == 1 {
// 		return 1
// 	}

// 	if n == 1 {
// 		return m
// 	}

// 	i := 2

// 	for {
// 		res := 1
// 		for j := 0; j < n; j++ {
// 			res = res * i
// 		}

// 		if res == m {
// 			return i
// 		}

// 		if res > m {
// 			return -1
// 		}

// 		i++
// 	}
// }

func main() {
	fmt.Println(nthRoot(1, 1))
	fmt.Println(nthRoot(3, 27))
	fmt.Println(nthRoot(4, 69))
	fmt.Println(nthRoot(2, 64))
}
