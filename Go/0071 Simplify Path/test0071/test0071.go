package test0071

import (
	"fmt"
	"time"
)

type testFunc func(string) string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("/home/", "/home")
	equals("/../", "/")
	equals("/home//foo/", "/home/foo")
	equals("/a/./b/../../c/", "/c")
	equals("/a/../../b/../c//.//", "/c")
	equals("/a//b////c/d//././/..", "/a/b/c")

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n string, expected string) {
	numTests++
	got := testFunction(n)
	if got != expected {
		fmt.Println(n, "got:", got, "expected:", expected)
		failed++
	}
}
