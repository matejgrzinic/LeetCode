package test0215

import (
	"fmt"
	"time"
)

type testFunc func([]int, int) int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{3, 2, 1, 5, 6, 4}, 2, 5)
	equals([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, target int, expected int) {
	numTests++
	got := testFunction(nums, target)
	if got != expected {
		fmt.Println(nums, target, "got:", got, "expected:", expected)
		failed++
	}
}
