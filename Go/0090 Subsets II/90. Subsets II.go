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
		combination(nums, &result, 0)
	} else {
		return [][]int{{}}
	}
	// result = append(result, []int{})
	noDuplicates := [][]int{{}}
	for i := 0; i < len(result); i++ {
		unique := true
		for j := 0; j < len(noDuplicates); j++ {
			if sliceEquals(result[i], noDuplicates[j]) {
				unique = false
				break
			}
		}
		if unique {
			noDuplicates = append(noDuplicates, result[i])
		}
	}
	return noDuplicates
}

func combination(nums []int, result *[][]int, index int) {
	if index == len(nums) {
		return
	}
	dup := false
	if index > 0 && nums[index] == nums[index-1] {
		dup = true
	}
	newResults := [][]int{}
	for i := 0; i < len(*result); i++ {
		// if dup && (*result)[i][len((*result)[i])-1] <= nums[index] {
		// 	continue
		// }
		tmp := make([]int, len((*result)[i]))
		copy(tmp, (*result)[i])
		tmp = append(tmp, nums[index])
		newResults = append(newResults, tmp)
	}
	if !dup {
		newResults = append(newResults, []int{nums[index]})
	}
	*result = append(*result, newResults...)
	combination(nums, result, index+1)
}

func sliceEquals(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
