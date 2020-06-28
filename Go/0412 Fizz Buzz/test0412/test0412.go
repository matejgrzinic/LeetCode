package test0412

import (
	"fmt"
	"time"
)

type testFunc func(int) []string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(1, []string{"1"})
	equals(2, []string{"1", "2"})
	equals(3, []string{"1", "2", "Fizz"})
	equals(15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n int, expected []string) {
	numTests++
	got := testFunction(n)
	if !sliceEquals(got, expected) {
		fmt.Println(n, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 []string, slice2 []string) bool {
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
