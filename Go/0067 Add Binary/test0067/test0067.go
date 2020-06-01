package test0067

import (
	"fmt"
	"time"
)

type testFunc func(string, string) string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("11", "1", "100")
	equals("1010", "1011", "10101")

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n string, k string, expected string) {
	numTests++
	got := testFunction(n, k)
	if got != expected {
		fmt.Println(n, k, "got:", got, "expected:", expected)
		failed++
	}
}
