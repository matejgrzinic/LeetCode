package main

import "./test0118"

func main() {
	test0118.Test(generate)
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	values := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		newRow := []int{1}
		for j := 1; j < len(values[i-1]); j++ {
			newRow = append(newRow, values[i-1][j-1]+values[i-1][j])
		}
		newRow = append(newRow, 1)
		values = append(values, newRow)
	}
	return values
}
