package main

func maxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}
