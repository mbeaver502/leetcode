// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/submissions/
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing
// a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction.
// If you cannot achieve any profit, return 0.
package main

import "log"

type example struct {
	prices []int
}

func main() {
	examples := []example{
		{[]int{7, 1, 5, 3, 6, 4}},
		{[]int{7, 6, 4, 3, 1}},
		{[]int{2, 4, 1}},
		{[]int{3, 2, 6, 5, 0, 3}},
		{[]int{4, 1, 5, 2, 7}},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v\n", ex.prices)
		log.Printf("output: %d\n", maxProfit(ex.prices))
	}
}

func maxProfit(prices []int) int {
	maxProfit := 0
	candidate := prices[0]
	for i := 1; i < len(prices); i++ {
		candidate = min(candidate, prices[i])
		maxProfit = max(maxProfit, prices[i]-candidate)
	}
	return maxProfit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
