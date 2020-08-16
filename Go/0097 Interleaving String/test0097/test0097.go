package test0097

import (
	"fmt"
	"time"
)

type testFunc func(string, string, string) bool

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("db", "b", "cbb", false)
	equals("aabcc", "", "aabcc", true)
	equals("aabcc", "dbbca", "aadbbcbcac", true)
	equals("aabcc", "dbbca", "aadbbbaccc", false)
	equals("", "", "", true)
	equals("a", "b", "ab", true)
	equals("a", "b", "ba", true)
	equals("aa", "b", "ba", false)
	equals("ba", "b", "ba", false)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n string, k string, l string, expected bool) {
	numTests++
	got := testFunction(n, k, l)
	if got != expected {
		fmt.Println(n, k, l, "got:", got, "expected:", expected)
		failed++
	}
}
