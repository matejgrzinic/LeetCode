package test####

import (
	"fmt"
	"time"
)

type testFunc func(string, string) bool
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

func equals(n string, k string, expected bool) {
	numTests++
	got := testFunction(n, k)
	if got != expected {
		fmt.Println(n, k, "got:", got, "expected:", expected)
		failed++
	}
}