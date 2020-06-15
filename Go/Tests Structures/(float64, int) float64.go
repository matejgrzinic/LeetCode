package test####

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