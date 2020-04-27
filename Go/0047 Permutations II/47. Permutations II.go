package main

import (
	"sort"

	"./test0047"
)

func main() {
	test0047.Test(permuteUnique)
}

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	copyAppend(&result, &nums)
	for {
		if nextPermutation(nums) {
			copyAppend(&result, &nums)
		} else {
			break
		}
	}
	return result
}

func copyAppend(result *[][]int, nums *[]int) {
	copyarray := make([]int, len(*nums))
	copy(copyarray, (*nums)[:])
	*result = append(*result, copyarray)
}

func nextPermutation(nums []int) bool {
	x := -1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			x = i - 1
		}
	}
	if x == -1 {
		return false
	}

	y := -1
	for i := 0; i < len(nums); i++ {
		if nums[x] < nums[i] {
			y = i
		}
	}
	nums[x], nums[y] = nums[y], nums[x]

	for i, j := x+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return true
}
