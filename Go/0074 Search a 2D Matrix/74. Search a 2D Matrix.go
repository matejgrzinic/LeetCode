package main

import (
	"./test0074"
)

func main() {
	test0074.Test(searchMatrix)
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 || target < matrix[0][0] {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	
	return binary(matrix, target, 0, (n*m)-1)
}

func binary(matrix [][]int, target int, left int, right int) bool {
	mid := (right + left) / 2
	i, j := mid/len(matrix[0]), mid%len(matrix[0])
	
	if matrix[i][j] == target {
		return true
	}
	if left >= right {
		return false
	}

	if target > matrix[i][j] {
		return binary(matrix, target, mid+1, right)
	}
	return binary(matrix, target, left, mid)
}
