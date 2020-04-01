package main

import (
	"sort"

	"./test0040"
)

func main() {
	test0040.Test(combinationSum2)
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	for i := 0; i < len(candidates); i++ {
		oneCombination(candidates[i+1:], target-candidates[i], []int{candidates[i]}, &result)
	}
	return result
}

func oneCombination(candidates []int, target int, current []int, result *[][]int) {
	if target == 0 {

		for i := 0; i < len(*result); i++ {
			if len((*result)[i]) == len(current) {
				same := true
				for j := 0; j < len((*result)[i]); j++ {
					if (*result)[i][j] != current[j] {
						same = false
						break
					}
				}
				if same {
					return
				}
			}
		}
		copyarray := make([]int, len(current))
		copy(copyarray, current[:])
		*result = append(*result, copyarray)

	}
	if target < 0 {
		return
	}
	if len(candidates) > 0 {
		oneCombination(candidates[1:], target, current, result)
		current = append(current, candidates[0])
		oneCombination(candidates[1:], target-candidates[0], current, result)
	}
}
