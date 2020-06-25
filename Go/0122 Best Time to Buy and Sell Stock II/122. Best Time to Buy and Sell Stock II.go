package main

import (
	"./test0122"
)

func main(){
	test0122.Test(maxProfit)
}

func maxProfit(prices []int) int {
	profit := 0	
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1]{
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}