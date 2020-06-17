package main

import (
	"fmt"

	"./test0037"
)

func main() {
	test0037.Test(solveSudoku)
}

func solveSudoku(board [][]byte) {
	rows := [9][9]int{}
	columns := [9][9]int{}
	squares := [9][9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 46 {
				num := board[i][j] - 49
				rows[i][num] = 1
				columns[j][num] = 1
				squareIndex := i/3*3 + j/3
				squares[squareIndex][num] = 1
			}
		}
	}

	backtracking(&board, &rows, &columns, &squares, 0, 0)
}

func backtracking(board *[][]byte, rows *[9][9]int, columns *[9][9]int, squares *[9][9]int, y int, x int) bool {
	if x == 9 {
		x = 0
		y++
	}
	if y == 9 {
		return true
	}

	if (*board)[y][x] == 46 {
		for i := 0; i < 9; i++ {
			if rows[y][i] == 0 && columns[x][i] == 0 && squares[y/3*3+x/3][i] == 0 {
				rows[y][i] = 1
				columns[x][i] = 1
				squares[y/3*3+x/3][i] = 1
				(*board)[y][x] = byte(49 + i)

				if backtracking(board, rows, columns, squares, y, x+1) {
					return true
				}

				rows[y][i] = 0
				columns[x][i] = 0
				squares[y/3*3+x/3][i] = 0
				(*board)[y][x] = byte(46)
			}
		}
	} else {
		if backtracking(board, rows, columns, squares, y, x+1) {
			return true
		}
	}
	return false
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
