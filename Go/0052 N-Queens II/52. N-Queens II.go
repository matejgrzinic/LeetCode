package main

import (
	"./test0052"
)

func main() {
	test0052.Test(totalNQueens)
}

func totalNQueens(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += helper([]int{i}, n)
	}
	return result
}

func helper(positions []int, n int) int {
	result := 0
	if len(positions) == n {
		return 1
	}

	for i := 0; i < n; i++ {
		isValid := true
		for j := 0; j < len(positions); j++ {
			if positions[j] == i || isDiagonal(j, positions[j], len(positions), i) {
				isValid = false
				break
			}
		}
		if isValid {
			copyarray := make([]int, len(positions))
			copy(copyarray, positions)
			copyarray = append(copyarray, i)
			result += helper(copyarray, n)
		}
	}
	return result
}

func isDiagonal(y1 int, x1 int, y2 int, x2 int) bool {
	a := absoluteInt(y1 - y2)
	b := absoluteInt(x1 - x2)
	if a == b {
		return true
	}
	return false
}

func absoluteInt(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
