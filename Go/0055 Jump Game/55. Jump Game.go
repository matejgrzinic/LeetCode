package main

import (
	"./test0055"
)

func main() {
	test0055.Test(canJump)
}

func canJump(nums []int) bool {
	t := make([]bool, len(nums))
	t[0] = true
	max := 0
	for i := 0; i < len(nums); i++ {
		if t[i] {
			l := nums[i] + i
			if l >= len(nums) {
				return true
			}
			for max < l {
				max++
				t[max] = true
			}
		} else {
			return false
		}
	}
	return true
}
