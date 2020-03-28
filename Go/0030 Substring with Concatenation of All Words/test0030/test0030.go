package test0030

import (
	"fmt"
	"time"
)

type testFunc func(string, []string) []int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("barfoothefoobarman", []string{"bar", "foo"}, []int{0, 9})
	equals("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{})
	equals("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(s string, words []string, expected []int) {
	numTests++
	got := testFunction(s, words)
	if !sliceEquals(got, expected) {
		fmt.Println(s, words, "got:", got, "expected:", expected)
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
