package main

import (
	"./test0073"
)

func main() {
	test0073.Test(setZeroes)
}

func setZeroes(matrix [][]int) {
	rows := make([]int, len(matrix))
	columns := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				columns[j] = 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] == 1 || columns[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
