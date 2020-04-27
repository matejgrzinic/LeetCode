package test0048

import (
	"fmt"
	"time"
)

type testFunc func([][]int)

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})
	equals([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}})
	equals([][]int{{2, 29, 20, 26, 16, 28}, {12, 27, 9, 25, 13, 21}, {32, 33, 32, 2, 28, 14}, {13, 14, 32, 27, 22, 26}, {33, 1, 20, 7, 21, 7}, {4, 24, 1, 6, 32, 34}}, [][]int{{4, 33, 13, 32, 12, 2}, {24, 1, 14, 33, 27, 29}, {1, 20, 32, 32, 9, 20}, {6, 7, 27, 2, 25, 26}, {32, 21, 22, 28, 13, 16}, {34, 7, 26, 14, 21, 28}})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(board [][]int, expected [][]int) {
	numTests++
	testFunction(board)
	if !sliceEquals(board, expected) {
		fmt.Println("got:", board, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 [][]int, slice2 [][]int) bool {
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
