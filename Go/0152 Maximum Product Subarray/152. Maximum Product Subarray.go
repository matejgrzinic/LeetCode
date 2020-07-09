package main

import "./test0152"

func main(){
	test0152.Test(maxProduct)
}

func maxProduct(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	a, b := nums[0], nums[0]
	result := nums[0]
	for _, num := range nums[1:] {
		a, b = max(max(a*num, b*num), num), min(min(a*num, b*num), num)
		if a > result{
			result = a
		}
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}