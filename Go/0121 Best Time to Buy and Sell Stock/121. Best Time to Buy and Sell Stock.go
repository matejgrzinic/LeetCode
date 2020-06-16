package main

import "./test0121"

func main() {
	test0121.Test(maxProfit)
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	m := 0
	l, r := prices[0], prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < l {
			l, r = prices[i], prices[i]
		} else if prices[i] > r {
			r = prices[i]
			if r-l > m {
				m = r - l
			}
		}
	}
	return m
}
