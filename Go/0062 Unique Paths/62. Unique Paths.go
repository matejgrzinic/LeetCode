package main

import (
	"./test0062"
)

func main() {
	test0062.Test(uniquePaths)
}

func uniquePaths2(m int, n int) int {
	paths := 0
	if m == 1 && n == 1 {
		return 1
	}
	if m > 1 {
		paths += uniquePaths(m-1, n)
	}
	if n > 1 {
		paths += uniquePaths(m, n-1)
	}
	return paths
}

func uniquePaths(m int, n int) int {
	seen := [][]int{}
	for i := 0; i < n; i++ {
		seen = append(seen, make([]int, m))
	}
	seen[n-1][m-1] = 1
	return nextStep(0, 0, seen)
}

func nextStep(y int, x int, seen [][]int) int {
	if seen[y][x] > 0 {
		return seen[y][x]
	}
	paths := 0
	if y < len(seen)-1 {
		paths += nextStep(y+1, x, seen)
	}
	if x < len(seen[0])-1 {
		paths += nextStep(y, x+1, seen)
	}
	seen[y][x] = paths
	return paths
}
