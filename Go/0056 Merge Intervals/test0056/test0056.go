package test0056

import (
	"fmt"
	"time"
)

type testFunc func([][]int) [][]int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, [][]int{{1, 10}})
	equals([][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}})
	equals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}})
	equals([][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}})
	equals([][]int{{1, 4}, {1, 4}}, [][]int{{1, 4}})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums [][]int, expected [][]int) {
	numTests++
	got := testFunction(nums)
	if !sliceEquals(got, expected) {
		fmt.Println(nums, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 [][]int, slice2 [][]int) bool {
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
