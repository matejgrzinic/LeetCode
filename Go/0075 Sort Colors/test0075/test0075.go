package test0075

import (
	"fmt"
	"time"
)

type testFunc func([]int)

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2})
	equals([]int{2, 0, 1}, []int{0, 1, 2})
	equals([]int{}, []int{})
	equals([]int{1}, []int{1})
	equals([]int{0}, []int{0})
	equals([]int{2}, []int{2})
	equals([]int{1, 2}, []int{1, 2})
	equals([]int{2, 1}, []int{1, 2})
	equals([]int{0, 2}, []int{0, 2})
	equals([]int{2, 0}, []int{0, 2})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(board []int, expected []int) {
	numTests++
	dup := duplicate(board)
	testFunction(board)
	if !sliceEquals(board, expected) {
		fmt.Println(dup, "got:", board, "expected:", expected)
		failed++
	}
}

func duplicate(matrix []int) []int {
	duplicate := make([]int, len(matrix))
	copy(duplicate, matrix)
	return duplicate
}

func sliceEquals(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
