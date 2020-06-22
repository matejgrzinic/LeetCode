package test0080

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

	equals([]int{1, 2}, 2)
	equals([]int{1, 1, 1, 2, 2, 3}, 5)
	equals([]int{}, 0)
	equals([]int{1}, 1)
	equals([]int{1, 1}, 2)
	equals([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7)

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
