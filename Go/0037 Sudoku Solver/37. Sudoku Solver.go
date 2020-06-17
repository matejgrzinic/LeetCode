package main

import (
	"fmt"

	"./test0037"
)

func main() {
	test0037.Test(solveSudoku)
}

func solveSudoku(board [][]byte) {
	makeChange(board)
}

func makeChange(board [][]byte) bool {
	y, x := findEmpty(board)
	if y == -1 {
		return true
	}
	for i := 1; i <= 9; i++ {
		board[y][x] = byte(48 + i)
		if checkRow(board, y) && checkColumn(board, x) && checkSquare(board, y, x) {
			if makeChange(board) {
				return true
			}
		}
	}
	board[y][x] = 46
	return false
}

func findEmpty(board [][]byte) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) == "." {
				return i, j
			}
		}
	}
	return -1, -1
}

func checkRow(board [][]byte, row int) bool {
	m := map[byte]bool{}
	for i := 0; i < 9; i++ {
		if string(board[row][i]) != "." {
			if _, ok := m[board[row][i]]; ok {
				//fmt.Println("Row")
				return false
			}
			m[board[row][i]] = true
		}
	}
	return true
}

func checkColumn(board [][]byte, column int) bool {
	m := map[byte]bool{}
	for i := 0; i < 9; i++ {
		if string(board[i][column]) != "." {
			if _, ok := m[board[i][column]]; ok {
				//fmt.Println("Column")
				return false
			}
			m[board[i][column]] = true
		}
	}
	return true
}

func checkSquare(board [][]byte, row int, column int) bool {
	m := map[byte]bool{}
	y := row / 3 * 3
	x := column / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if string(board[y+i][x+j]) != "." {
				if _, ok := m[board[y+i][x+j]]; ok {
					//fmt.Println("Square")
					return false
				}
				m[board[y+i][x+j]] = true
			}
		}
	}
	return true
}

func print(board [][]byte) {
	for i := 0; i < len(board); i++ {
		rowText := ""
		for j := 0; j < len(board[0]); j++ {
			rowText += string(board[i][j]) + " "
		}
		fmt.Println(rowText)
	}
	fmt.Println()
}
