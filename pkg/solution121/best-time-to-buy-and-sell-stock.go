package solution121

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	var maxProfit int

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
