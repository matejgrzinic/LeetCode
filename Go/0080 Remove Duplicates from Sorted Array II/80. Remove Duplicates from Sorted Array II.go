package main

import (
	"./test0080"
)

func main() {
	test0080.Test(removeDuplicates)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	last, d, newIndex := nums[0], false, 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == last {
			if !d {
				d = true
				nums[i], nums[newIndex] = nums[newIndex], nums[i]
				newIndex++
			}
		} else {
			last = num
			d = false
			nums[i], nums[newIndex] = nums[newIndex], nums[i]
			newIndex++
		}
	}
	return newIndex
}
