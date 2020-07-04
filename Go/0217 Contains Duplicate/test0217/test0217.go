package test0217

import (
	"fmt"
	"time"
)

type testFunc func([]int) bool

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{1, 2, 3, 1}, true)
	equals([]int{1, 2, 3, 4}, false)
	equals([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, expected bool) {
	numTests++
	got := testFunction(nums)
	if got != expected {
		fmt.Println(nums, "got:", got, "expected:", expected)
		failed++
	}
}
