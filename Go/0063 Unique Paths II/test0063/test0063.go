package test0063

import (
	"fmt"
	"time"
)

type testFunc func([][]int) int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(board [][]int, expected int) {
	numTests++
	got := testFunction(board)
	if got != expected {
		fmt.Println(board, "got:", got, "expected:", expected)
		failed++
	}
}
