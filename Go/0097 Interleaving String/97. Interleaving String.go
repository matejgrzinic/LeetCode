package main

import (
	"./test0097"
)

func main() {
	test0097.Test(isInterleave)
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	matrix := make([][]bool, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]bool, len(s2)+1)
	}

	matrix[0][0] = true

	for i := 1; i < len(matrix); i++ {
		if s1[i-1] == s3[i-1] && matrix[i-1][0] {
			matrix[i][0] = true
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if s2[i-1] == s3[i-1] && matrix[0][i-1] {
			matrix[0][i] = true
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			index := i + j - 1
			if s1[i-1] == s3[index] && matrix[i-1][j] {
				matrix[i][j] = true
			} else if s2[j-1] == s3[index] && matrix[i][j-1] {
				matrix[i][j] = true
			}
		}
	}

	return matrix[len(matrix)-1][len(matrix[0])-1]
}
