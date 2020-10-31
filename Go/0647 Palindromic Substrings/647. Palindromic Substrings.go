package main

func countSubstrings(s string) int {
	matrix := make([][]bool, len(s))
	for i := range matrix {
		matrix[i] = make([]bool, len(s))
	}
	result := len(s)
	for i := 1; i < len(s); i++ {
		for j := 0; j < i-1; j++ {
			if matrix[i-1][j+1] && s[i] == s[j] {
				matrix[i][j] = true
				result++
			}
		}
		if s[i-1] == s[i] {
			matrix[i][i-1] = true
			result++
		}
		matrix[i][i] = true
	}
	return result
}
