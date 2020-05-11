package test0062

import (
	"fmt"
	"time"
)

type testFunc func(int, int) int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(3, 2, 3)
	equals(7, 3, 28)
	equals(19, 13, 86493225)
	equals(50, 50, 858110510779117752)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(m int, n int, expected int) {
	numTests++
	got := testFunction(m, n)
	if got != expected {
		fmt.Println(m, n, "got:", got, "expected:", expected)
		failed++
	}
}
