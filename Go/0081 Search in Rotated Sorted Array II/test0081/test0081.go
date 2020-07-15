package test0081

import (
	"fmt"
	"time"
)

type testFunc func([]int, int) bool

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{1, 2, 3, 4, 5}, 2, true)
	equals([]int{3, 1}, 1, true)
	equals([]int{3, 1, 2, 2, 2}, 1, true)
	equals([]int{1, 3, 1, 1, 1}, 3, true)
	equals([]int{1, 1, 1, 1, 3}, 3, true)
	equals([]int{1, 3}, 3, true)
	equals([]int{5, 1, 3}, 3, true)
	equals([]int{3, 1, 1}, 3, true)
	equals([]int{}, 5, false)
	equals([]int{1}, 5, false)
	equals([]int{5}, 5, true)
	equals([]int{2, 5, 6, 0, 0, 1, 2}, 3, false)
	equals([]int{2, 5, 6, 0, 0, 1, 2}, 3, false)
	equals([]int{1, 2, 3, 4, 5}, 1, true)

	equals([]int{1, 2, 3, 4, 5}, 3, true)
	equals([]int{1, 2, 3, 4, 5}, 4, true)
	equals([]int{1, 2, 3, 4, 5}, 5, true)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, target int, expected bool) {
	numTests++
	got := testFunction(nums, target)
	if got != expected {
		fmt.Println(nums, target, "got:", got, "expected:", expected)
		failed++
	}
}
