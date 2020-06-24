package main

import "./test0119"

func main() {
	test0119.Test(getRow)
}

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[0] = 1
		last := 1
		for j := 1; j <= i; j++ {
			last, result[j] = result[j], result[j]+last
		}
	}
	return result
}
