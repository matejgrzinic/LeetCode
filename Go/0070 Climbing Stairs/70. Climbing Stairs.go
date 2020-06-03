package main

import "./test0070"

func main() {
	test0070.Test(climbStairs)
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	x, y := 1, 2
	for i := 2; i < n; i++ {
		x, y = y, x+y
	}
	return y
}
