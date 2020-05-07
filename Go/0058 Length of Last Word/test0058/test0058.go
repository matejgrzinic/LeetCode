package test0058

import (
	"fmt"
	"time"
)

type testFunc func(string) int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("Hello World", 5)
	equals("a ", 1)
	equals(" ", 0)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n string, expected int) {
	numTests++
	got := testFunction(n)
	if got != expected {
		fmt.Println(n, "got:", got, "expected:", expected)
		failed++
	}
}
