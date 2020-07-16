package main

import (
	"./test0079"
)

func main() {
	test0079.Test(exist)
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] && step(&board, word[1:], [][]int{}, i, j) {
				return true
			}
		}
	}
	return false
}

func step(board *[][]byte, word string, visited [][]int, y int, x int) bool {
	if len(word) == 0 {
		return true
	}

	visited = append(visited, []int{y, x})

	if checkIfValid(board, word, visited, y+1, x) || checkIfValid(board, word, visited, y-1, x) ||
		checkIfValid(board, word, visited, y, x+1) || checkIfValid(board, word, visited, y, x-1) {
		return true
	}
	return false
}

func checkIfValid(board *[][]byte, word string, visited [][]int, y int, x int) bool {
	if y < len(*board) && y >= 0 && x >= 0 && x < len((*board)[0]) && (*board)[y][x] == word[0] {
		valid := true
		for _, vals := range visited {
			if y == vals[0] && x == vals[1] {
				valid = false
				break
			}
		}
		if valid {
			if step(board, word[1:], visited, y, x) {
				return true
			}
		}
	}
	return false
}
