package test0042

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

	equals([]int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3}, 83)
	equals([]int{5, 0, 5}, 5)
	equals([]int{5, 0, 0, 5}, 10)
	equals([]int{1, 2, 3, 4, 5}, 0)
	equals([]int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, 23)
	equals([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
	equals([]int{4, 2, 3}, 1)
	equals([]int{}, 0)

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
