package main

import (
	"fmt"

	"./test0045"
)

func main() {
	test0045.Test(jump)
}

func jump(nums []int) int {
	n := len(nums) - 1
	if n == 0 {
		return 0
	}
	jumps := make([]int, len(nums))
	for i := nums[0]; i > 0; i-- {
		if i >= n {
			return 1
		}
		jumps[i]++
	}
	for i := 1; i < n; i++ {
		rng := nums[i] + i
		if rng >= n {
			return jumps[i] + 1
		}
		for j := rng; true; j-- {
			if jumps[j] != 0 {
				break
			}
			jumps[j] = jumps[i] + 1
		}
	}
	fmt.Println(jumps)
	return 0
}
