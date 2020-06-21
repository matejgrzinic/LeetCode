package main

import (
	"./test0041"
)

func main() {
	test0041.Test(firstMissingPositive)
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val < 1 || val > len(nums) {
			continue
		}
		if nums[i] != nums[val-1] {
			nums[i], nums[val-1] = nums[val-1], nums[i]
			i--
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
