package test0038

import (
	"fmt"
	"time"
)

type testFunc func(int) string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(1, "1")
	equals(2, "11")
	equals(3, "21")
	equals(4, "1211")
	equals(5, "111221")
	equals(6, "312211")
	equals(7, "13112221")
	equals(8, "1113213211")
	equals(9, "31131211131221")
	equals(10, "13211311123113112211")

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n int, expected string) {
	numTests++
	got := testFunction(n)
	if got != expected {
		fmt.Println(n, "got:", got, "expected:", expected)
		failed++
	}
}
