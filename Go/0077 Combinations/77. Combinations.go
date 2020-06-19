package main

import (
	"./test0077"
)

func main() {
	test0077.Test(combine)
}

func combine(n int, k int) [][]int {
	res := [][]int{}
	permut(0, n, k, &res, []int{})
	return res
}

func permut(num int, n int, k int, res *[][]int, current []int) {
	if k == 0 {
		*res = append(*res, current)
	} else {
		for i := num + 1; i <= n-k+1; i++ {
			newSlice := make([]int, len(current))
			copy(newSlice, current)
			newSlice = append(newSlice, i)
			permut(i, n, k-1, res, newSlice)
		}
	}
}
