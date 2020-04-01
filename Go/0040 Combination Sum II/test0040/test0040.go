package test0040

import (
	"fmt"
	"time"
)

type testFunc func([]int, int) [][]int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{3, 2, 1, 1, 1}, 3, [][]int{{3}, {1, 2}, {1, 1, 1}})
	equals([]int{2, 1}, 2, [][]int{{2}})
	equals([]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}})
	equals([]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, target int, expected [][]int) {
	numTests++
	got := testFunction(nums, target)
	if !sliceEquals(got, expected) {
		fmt.Println(nums, target, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 [][]int, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {

		found := false
		for k := 0; k < len(slice2); k++ {
			if len(slice1[i]) != len(slice2[k]) {
				continue
			}
			same := true
			for j := 0; j < len(slice1[i]); j++ {
				if slice1[i][j] != slice2[k][j] {
					same = false
				}
			}
			if same {
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
