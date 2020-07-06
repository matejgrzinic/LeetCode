package test0179

import (
	"fmt"
	"time"
)

type testFunc func([]int) string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{10}, "10")
	equals([]int{10, 2}, "210")
	equals([]int{3, 30, 34, 5, 9}, "9534330")
	equals([]int{0, 0}, "0")
	equals([]int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}, "9609938824824769735703560743981399")

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, expected string) {
	numTests++
	got := testFunction(nums)
	if got != expected {
		fmt.Println(nums, "got:", got, "expected:", expected)
		failed++
	}
}
