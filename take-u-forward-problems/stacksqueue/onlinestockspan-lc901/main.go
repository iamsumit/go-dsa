package main

import "fmt"

// https://leetcode.com/problems/online-stock-span/

// Optimal approach
type Stock struct {
	price int
	span  int
}

type StockSpanner struct {
	data []Stock
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (s *StockSpanner) Next(price int) int {
	span := 1

	for len(s.data) > 0 && s.data[len(s.data)-1].price <= price {
		span += s.data[len(s.data)-1].span
		s.data = s.data[:len(s.data)-1]
	}

	s.data = append(s.data, Stock{
		price: price,
		span:  span,
	})

	return span
}

// // Brute force
// type StockSpanner struct {
// 	data []int
// }

// func Constructor() StockSpanner {
// 	return StockSpanner{}
// }

// func (s *StockSpanner) Next(price int) int {
// 	count := 1
// 	for i := len(s.data) - 1; i >= 0; i-- {
// 		if s.data[i] > price {
// 			break
// 		}

// 		count++
// 	}

// 	s.data = append(s.data, price)
// 	return count
// }

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func main() {
	obj := Constructor()
	fmt.Println("obj.Next(100)", obj.Next(100))
	fmt.Println("obj.Next(80)", obj.Next(80))
	fmt.Println("obj.Next(60)", obj.Next(60))
	fmt.Println("obj.Next(70)", obj.Next(70))
	fmt.Println("obj.Next(60)", obj.Next(60))
	fmt.Println("obj.Next(75)", obj.Next(75))
	fmt.Println("obj.Next(85)", obj.Next(85))
}
