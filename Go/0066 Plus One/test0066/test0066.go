package test0066

import (
	"fmt"
	"time"
)

type testFunc func([]int) []int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{1, 2, 3}, []int{1, 2, 4})
	equals([]int{4, 3, 2, 1}, []int{4, 3, 2, 2})
	equals([]int{1, 2, 9}, []int{1, 3, 0})
	equals([]int{9, 9, 9}, []int{1, 0, 0, 0})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(matrix []int, expected []int) {
	numTests++
	got := testFunction(matrix)
	if !sliceEquals(got, expected) {
		fmt.Println(matrix, "got:", got, "expected:", expected)
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
