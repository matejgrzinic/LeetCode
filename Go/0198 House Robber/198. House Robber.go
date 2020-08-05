package main

import (
	"./test0198"
)

func main() {
	test0198.Test(rob)
}

func rob(nums []int) int {
	var x, y int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			x = max(x+nums[i], y)
		} else {
			y = max(y+nums[i], x)
		}
	}
	return max(x, y)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
