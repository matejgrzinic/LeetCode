package main

import (
	"./test0078"
)

func main() {
	test0078.Test(subsets)
}

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		subset(nums, 0, i, []int{}, &res)
	}
	return res
}

func subset(nums []int, index int, size int, current []int, res *[][]int) {
	if size == 0 {
		*res = append(*res, current)
	} else {
		for i := index + 1; i <= len(nums)-size+1; i++ {
			newSlice := make([]int, len(current))
			copy(newSlice, current)
			newSlice = append(newSlice, nums[i-1])
			subset(nums, i, size-1, newSlice, res)
		}
	}
}
