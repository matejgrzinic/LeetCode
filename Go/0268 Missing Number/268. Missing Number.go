package main

import (
	"./test0268"
)

func main() {
	test0268.Test(missingNumber)
}

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num != i && num < len(nums) {
			nums[i], nums[num] = nums[num], nums[i]
			i--
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
