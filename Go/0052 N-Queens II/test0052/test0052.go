package test0052

import (
	"fmt"
	"time"
)

type testFunc func(int) int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(1, 1)
	equals(2, 0)
	equals(3, 0)
	equals(4, 2)
	equals(5, 10)
	equals(6, 4)
	equals(7, 40)
	equals(8, 92)
	equals(9, 352)
	equals(10, 724)
	equals(11, 2680)
	equals(12, 14200)
	equals(13, 73712)
	// equals(14, 365596)
	// equals(15, 2279184)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n int, expected int) {
	numTests++
	got := testFunction(n)
	if got != expected {
		fmt.Println(n, "got:", got, "expected:", expected)
		failed++
	}
}
