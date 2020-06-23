package test0137

import (
	"fmt"
	"time"
)

type testFunc func([]int) int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{2, 3, 2, 2}, 3)
	equals([]int{2, 2, 3, 2}, 3)
	equals([]int{0, 1, 0, 1, 0, 1, 99}, 99)
	equals([]int{2, 2, 2, 3, 1, 1, 1}, 3)
	equals([]int{2, 2, 2, 3}, 3)
	equals([]int{3, 2, 2, 2}, 3)
	equals([]int{3, 2, 2, 2}, 3)
	equals([]int{1, 1, 2, 3, 2, 3, 2, 6, 1, 7, 7, 7, 2, 8, 8, 8}, 6)
	equals([]int{1, 1, 2, 3, 2, 3, 2, 6, 1, 7, 7, 2, 7, 8, 8, 8}, 6)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, expected int) {
	numTests++
	got := testFunction(nums)
	if got != expected {
		fmt.Println(nums, "got:", got, "expected:", expected)
		failed++
	}
}
