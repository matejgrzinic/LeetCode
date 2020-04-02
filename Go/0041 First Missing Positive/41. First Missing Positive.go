package main

import "./test0041"

func main(){
	test0041.Test(firstMissingPositive)
}

func firstMissingPositive(nums []int) int {
	smallest := 213213
	for _, val := range nums {
		if val > 0 && val < smallest{
			smallest = val
		} 
	}
    return smallest
}