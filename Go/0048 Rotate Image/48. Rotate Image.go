package main

import (
	"./test0048"
)

func main() {
	test0048.Test(rotate)
}

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	d := 0
	for (n+1)-2*d > 0 {
		for i := 0; i < n-(d*2); i++ {
			matrix[d][d+i], matrix[i+d][n-d] = matrix[i+d][n-d], matrix[d][d+i]
			matrix[n-d][n-i-d], matrix[n-i-d][d] = matrix[n-i-d][d], matrix[n-d][n-i-d]
			matrix[d][d+i], matrix[n-d][n-i-d] = matrix[n-d][n-i-d], matrix[d][d+i]
		}
		d++
	}
}
