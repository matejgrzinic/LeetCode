package test0093

import (
	"fmt"
	"time"
)

type testFunc func(string) []string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("25525511135", []string{"255.255.11.135", "255.255.111.35"})
	equals("010010", []string{"0.10.0.10", "0.100.1.0"})
	equals("1111", []string{"1.1.1.1"})
	equals("111", []string{})
	equals("", []string{})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n string, expected []string) {
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
		found := false
		for _, slice := range slice2 {
			if slice1[i] == slice {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
