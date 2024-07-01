package main

import "fmt"

// https://takeuforward.org/data-structure/find-minimum-number-of-coins/
func minCoins(v int) (int, []int) {
	count := 0
	coins := []int{}

	for v != 0 {
		var coin int
		switch {
		case v >= 1000:
			coin = 1000
		case v >= 500:
			coin = 500
		case v >= 100:
			coin = 100
		case v >= 50:
			coin = 50
		case v >= 20:
			coin = 20
		case v >= 10:
			coin = 10
		case v >= 5:
			coin = 5
		case v >= 2:
			coin = 2
		case v >= 1:
			coin = 1
		}

		count += 1
		v = v - coin
		coins = append(coins, coin)
	}

	return count, coins
}

func main() {
	fmt.Println(minCoins(70))
	fmt.Println(minCoins(15))
	fmt.Println(minCoins(121))
}
