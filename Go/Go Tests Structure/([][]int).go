package test####

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


	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(board [][]int, expected [][]int) {
	numTests++
	dup := duplicate(board)
	testFunction(board)
	if !sliceEquals(board, expected) {
		fmt.Println(dup, "got:", board, "expected:", expected)
		failed++
	}
}

func duplicate(matrix [][]int) [][]int{
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
    	duplicate[i] = make([]int, len(matrix[i]))
    	copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func sliceEquals(slice1 [][]int, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if len(slice1[i]) != len(slice2[i]){
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