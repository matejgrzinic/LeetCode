package test0034

import (
	"fmt"
	"time"
)

type testFunc func([]int, int) []int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4})
	equals([]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1})
	equals([]int{8, 8, 8, 8, 8}, 8, []int{0, 4})
	equals([]int{}, 8, []int{-1, -1})
	equals([]int{8}, 8, []int{0, 0})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, target int, expected []int) {
	numTests++
	got := testFunction(nums, target)
	if !sliceEquals(got, expected) {
		fmt.Println(nums, target, "got:", got, "expected:", expected)
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
