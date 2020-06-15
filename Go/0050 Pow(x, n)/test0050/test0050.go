package test0050

import (
	"fmt"
	"time"
)

type testFunc func(float64, int) float64

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(2, -2, 0.25)
	equals(2.00, 10, 1024.00)
	equals(2.1, 3, 9.26100)
	equals(1, 10, 1)
	equals(5, 0, 1)
	equals(5, -1, 0.2)
	equals(1, -2147483648, 1)
	equals(8.88023, 3, 700.28148)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(m float64, n int, expected float64) {
	numTests++
	got := testFunction(m, n)
	if got != expected {
		fmt.Println(m, n, "got:", got, "expected:", expected)
		failed++
	}
}
