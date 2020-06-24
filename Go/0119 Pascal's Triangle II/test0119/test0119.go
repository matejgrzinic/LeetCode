package test0119

import (
	"fmt"
	"time"
)

type testFunc func(int) []int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(0, []int{1})
	equals(1, []int{1, 1})
	equals(2, []int{1, 2, 1})
	equals(3, []int{1, 3, 3, 1})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n int, expected []int) {
	numTests++
	got := testFunction(n)
	if !sliceEquals(got, expected) {
		fmt.Println(n, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
