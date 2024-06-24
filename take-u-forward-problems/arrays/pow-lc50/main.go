package main

import (
	"fmt"
)

// Better approach using binary exponention
func myPow(x float64, n int) float64 {
	ans := 1.0
	nn := n

	if nn < 0 {
		nn = -1 * nn
	}

	for nn != 0 {
		if nn%2 != 0 {
			ans = ans * x
			nn = nn - 1
		} else {
			x = x * x
			nn = nn / 2
		}
	}

	if n < 0 {
		ans = 1 / ans
	}

	return ans
}

// func myPow(x float64, n int) float64 {
// 	if n == 0 || x == 1 {
// 		return 1
// 	}

// 	absn := int(math.Abs(float64(n)))

// 	powout := x

// 	stop := 0

// 	for i := 2; i <= absn; i = i * 2 {
// 		powout = float64(powout * powout)
// 		if i*2 > absn {
// 			stop = i
// 		}
// 	}

// 	if stop != 0 {
// 		for i := stop + 1; i <= absn; i++ {
// 			powout = float64(powout * x)
// 		}
// 	}

// 	if n < 0 {
// 		powout = 1 / powout
// 	}

// 	return powout
// }

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}
