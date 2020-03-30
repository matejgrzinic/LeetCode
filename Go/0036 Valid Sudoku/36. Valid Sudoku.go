package main

import (
	"./test0036"
)

func main() {
	test0036.Test(isValidSudoku)
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := make(map[string]bool)
		for j := 0; j < 9; j++ {
			ch := string(board[i][j])
			if _, ok := m[ch]; ok {
				if ch != "." {
					return false
				}
			} else {
				m[ch] = true
			}
		}
	}
	for i := 0; i < 9; i++ {
		m := make(map[string]bool)
		for j := 0; j < 9; j++ {
			ch := string(board[j][i])
			if _, ok := m[ch]; ok {
				if ch != "." {
					return false
				}
			} else {
				m[ch] = true
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m := make(map[string]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					ch := string(board[i*3+k][j*3+l])
					if _, ok := m[ch]; ok {
						if ch != "." {
							return false
						}
					} else {
						m[ch] = true
					}
				}
			}
		}
	}
	return true
}
