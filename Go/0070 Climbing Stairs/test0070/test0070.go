package test0070

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
	equals(2, 2)
	equals(3, 3)
	equals(15, 987)
	equals(50, 20365011074)
	equals(100, 1298777728820984005)

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
