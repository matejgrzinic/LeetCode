package test0131

import (
	"fmt"
	"time"
)

type testFunc func(string) [][]string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("fff", [][]string{{"f", "f", "f"}, {"f", "ff"}, {"ff", "f"}, {"fff"}})
	equals("cdd", [][]string{{"c", "d", "d"}, {"c", "dd"}})
	equals("aab", [][]string{{"aa", "b"}, {"a", "a", "b"}})
	equals("a", [][]string{{"a"}})
	equals("", [][]string{})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(words string, expected [][]string) {
	numTests++
	got := testFunction(words)
	if !sliceEquals(got, expected) {
		fmt.Println(words, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 [][]string, slice2 [][]string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		found := false
		for j := 0; j < len(slice2); j++ {
			if stringEquals(slice1[i], slice2[j]) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func stringEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
