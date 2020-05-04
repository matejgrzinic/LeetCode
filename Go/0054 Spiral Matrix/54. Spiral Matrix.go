package main

import (
	"./test0054"
)

func main() {
	test0054.Test(spiralOrder)
}

func spiralOrder(matrix [][]int) []int {
	path := []int{}
	if len(matrix) == 0 {
		return path
	}
	dir := []bool{true, true}
	speed := []int{len(matrix[0]), len(matrix) - 1}
	turnY := false
	position := []int{-1, 0}

	for {
		active := 0
		if turnY {
			active = 1
		}
		turnY = !turnY
		if speed[active] == 0 {
			break
		}
		step := 1
		if !dir[active] {
			step = -1
		}
		for i := 0; i < speed[active]; i++ {
			position[active] += step
			path = append(path, matrix[position[1]][position[0]])
		}
		dir[active] = !dir[active]
		speed[active]--
	}
	return path
}
