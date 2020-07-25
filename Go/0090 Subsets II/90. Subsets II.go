package main

import (
	"sort"

	"./test0090"
)

func main() {
	test0090.Test(subsetsWithDup)
}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	if len(nums) > 0 {
		sort.Ints(nums)
		combination(nums, &result, []int{}, 0)
		return result
	}
	return [][]int{{}}
}

func combination(nums []int, result *[][]int, current []int, index int) {
	tmp := make([]int, len(current))
	copy(tmp, current)
	*result = append(*result, tmp)

	for i := index; i < len(nums); i++ {
		if i != index && nums[i] == nums[i-1] {
			continue
		}
		current = append(current, nums[i])
		combination(nums, result, current, i+1)
		current = current[:len(current)-1]
	}
}
