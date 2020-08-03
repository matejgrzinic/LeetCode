package main

import (
	"./test0044"
)

func main() {
	test0044.Test(isMatch)
}

func isMatch(s string, p string) bool {
	dynamic := make([][]bool, len(s)+1)
	for i := 0; i < len(dynamic); i++ {
		dynamic[i] = make([]bool, len(p)+1)
	}

	dynamic[0][0] = true

	for i := 1; i < len(dynamic[0]); i++ {
		if dynamic[0][i-1] && p[i-1] == '*' {
			dynamic[0][i] = true
		}
	}

	for i := 1; i < len(dynamic); i++ {
		for j := 1; j < len(dynamic[0]); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dynamic[i][j] = dynamic[i-1][j-1]
			} else if p[j-1] == '*' {
				if dynamic[i-1][j] || dynamic[i][j-1] {
					dynamic[i][j] = true
				}
			}
		}
	}
	return dynamic[len(s)][len(p)]
}
