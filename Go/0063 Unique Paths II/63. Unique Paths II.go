package main

import (
	"./test0063"
)

func main() {
	test0063.Test(uniquePathsWithObstacles)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	matrix[n-1][m-1] = 1
	result := newStep(0, 0, obstacleGrid, matrix)
	//fmt.Println(matrix)
	return result
}

func newStep(y int, x int, obstacleGrid [][]int, matrix [][]int) int {
	if obstacleGrid[y][x] == 1 {
		return 0
	}

	if matrix[y][x] > 0 {
		return matrix[y][x]
	}

	result := 0
	if y < len(obstacleGrid)-1 {
		result += newStep(y+1, x, obstacleGrid, matrix)
	}
	if x < len(obstacleGrid[0])-1 {
		result += newStep(y, x+1, obstacleGrid, matrix)
	}
	matrix[y][x] = result
	return result
}
