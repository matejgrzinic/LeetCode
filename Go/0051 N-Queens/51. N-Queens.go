package main

import (
	"./test0051"
)

func main() {
	test0051.Test(solveNQueens)
}

func solveNQueens(n int) [][]string {
	result := [][]string{}
	for i := 0; i < n; i++ {
		helper([]int{i}, n, &result)
	}
	return result
}

func helper(positions []int, n int, result *[][]string) {
	if len(positions) == n {
		solution := []string{}
		for _, i := range positions {
			str := ""
			for j := 0; j < i; j++ {
				str += "."
			}
			str += "Q"
			for j := i + 1; j < n; j++ {
				str += "."
			}
			solution = append(solution, str)
		}
		*result = append(*result, solution)
		return
	}
	newY := len(positions)
	for i := 0; i < n; i++ {
		isValid := true

		for j := 0; j < len(positions); j++ {
			if positions[j] == i {
				isValid = false
				break
			}
			if isDiagonal(j, positions[j], newY, i) {
				isValid = false
				break
			}
		}

		if isValid {
			copyarray := make([]int, len(positions))
			copy(copyarray, positions)
			copyarray = append(copyarray, i)
			helper(copyarray, n, result)
		}
	}
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
