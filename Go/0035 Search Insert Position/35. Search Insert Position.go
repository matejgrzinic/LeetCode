package main

import "./test0035"

func main() {
	//test0035.Test(searchInsert)
	test0035.Test(searchInsert)
}

func searchInsert(nums []int, target int) int {
	for index := range nums {
		if target <= nums[index] {
			return index
		}
	}
	return len(nums)
}
