package main

import (
	"./test0059"
)

func main() {
	test0059.Test(generateMatrix)
}

func generateMatrix(n int) [][]int {
	speed := [2]int{n, n - 1}
	position := [2]int{-1, 0}
	direction := [2]bool{true, true}
	turn := true
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	counter := 1
	for true {
		index := 1
		if turn {
			index = 0
		}
		if speed[index] == 0 {
			return result
		}
		turn = !turn
		step := -1
		if direction[index] {
			step = 1
		}
		for i := 0; i < speed[index]; i++ {
			position[index] += step
			result[position[1]][position[0]] = counter
			counter++
		}
		direction[index] = !direction[index]
		speed[index]--
	}
	return result
}
