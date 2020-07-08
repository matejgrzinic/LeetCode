package main

import "./test0216"

func main() {
	test0216.Test(combinationSum3)
}

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	makeCombination(k, n, []int{}, &result)
	return result
}

func makeCombination(k int, n int, current []int, result *[][]int) {
	if k == 0 {
		if n == 0 {
			currCopy := make([]int, len(current))
			copy(currCopy, current)
			*result = append(*result, currCopy)
		}
		return
	}
	if n < 1 {
		return
	}
	last := 1
	if len(current) > 0 {
		last = current[len(current)-1] + 1
	}

	for i := last; i <= 9; i++ {
		current = append(current, i)
		makeCombination(k-1, n-i, current, result)
		current = current[:len(current)-1]
	}
}
