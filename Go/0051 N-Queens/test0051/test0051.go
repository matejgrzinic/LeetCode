package test0051

import (
	"fmt"
	"time"
)

type testFunc func(int) [][]string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(0, [][]string{})
	equals(1, [][]string{{"Q"}})
	equals(2, [][]string{})
	equals(3, [][]string{})
	equals(4, [][]string{{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."}})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n int, expected [][]string) {
	numTests++
	got := testFunction(n)
	if !sliceEquals(got, expected) {
		fmt.Println(n, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 [][]string, slice2 [][]string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}
		for j := 0; j < len(slice1[i]); j++ {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}
	}
	return true
}
