package main

import "fmt"

func maxProfit(prices []int) int {
	bestProfit := 0
	maxPrice := 0
	n := len(prices)
	for i := n - 2; i >= 0; i-- {

		if prices[i] > maxPrice {
			maxPrice = prices[i]
		} else {
			if maxPrice-prices[i] > bestProfit {
				bestProfit = maxPrice - prices[i]
			}
		}

	}
	return bestProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max Profit:", maxProfit(prices))
}
