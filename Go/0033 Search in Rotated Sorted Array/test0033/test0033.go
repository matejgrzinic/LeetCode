package test0033

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

	equals([]int{4, 5, 6, 7, 0, 1, 2}, 4, 0)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 5, 1)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 6, 2)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 7, 3)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 0, 4)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 1, 5)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 2, 6)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 3, -1)
	equals([]int{5, 1, 3}, 2, -1)
	equals([]int{8, 1, 3}, 8, 0)
	equals([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4)
	equals([]int{3, 1}, 1, 1)
	equals([]int{5, 1, 3}, 3, 2)
	equals([]int{1, 3}, 3, 1)
	equals([]int{5, 1, 2, 3, 4}, 1, 1)

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
