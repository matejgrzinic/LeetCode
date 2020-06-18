package test0077

import (
	"fmt"
	"time"
)

type testFunc func(int, int) [][]int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(2, 1, [][]int{{1}, {2}})
	equals(4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}})
	equals(1, 1, [][]int{{1}})
	equals(3, 3, [][]int{{1, 2, 3}})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums int, target int, expected [][]int) {
	numTests++
	got := testFunction(nums, target)
	if !sliceEqualsDifferentOrder(got, expected) {
		fmt.Println(nums, target, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEqualsDifferentOrder(slice1 [][]int, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		found := false
		for k := 0; k < len(slice2); k++ {
			if len(slice1[i]) == len(slice2[k]) {
				found = true
				for j := 0; j < len(slice2[k]); j++ {
					if slice1[i][j] != slice2[k][j] {
						found = false
						break
					}
				}
			}
			if found {
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
