package main

import (
	"./test0056"
)

func main() {
	test0056.Test(merge)
}

func merge(intervals [][]int) [][]int {
	results := [][]int{}
	for _, interval := range intervals {
		a, b := interval[0], interval[1]
		insertNew := true
		for i, result := range results {
			c, d := result[0], result[1]

			if a <= d && b > d {
				results[i][1] = b
				if a < c {
					results[i][0] = a
				}
				insertNew = false
				break
			} else if a < c && b >= c {
				results[i][0] = a
				if b > d {
					results[i][1] = b
				}
				insertNew = false
				break
			} else if a >= c && b <= d {
				insertNew = false
				break
			}
		}
		if insertNew {
			results = append(results, []int{a, b})
		}
	}

	if !intervalEquals(intervals, results) {
		results = merge(results)
	}
	return results
}

func intervalEquals(slice1 [][]int, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}
		for j := 0; j < len(slice1[i]); j++ {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}
	}
	return true
}
