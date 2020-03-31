package test0039

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

	equals([]int{2, 1}, 2, [][]int{{2}, {1, 1}})
	equals([]int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}})
	equals([]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}})
	equals([]int{2, 5, 8, 4}, 10, [][]int{{5, 5}, {2, 4, 4}, {2, 8}, {2, 2, 2, 4}, {2, 2, 2, 2, 2}})
	equals([]int{2, 4}, 10, [][]int{{2, 2, 2, 2, 2}, {2, 4, 4}, {2, 2, 2, 4}})
	equals([]int{5, 10, 8, 4, 3, 12, 9}, 27, [][]int{{3, 3, 3, 3, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 3, 3, 4, 5}, {3, 3, 3, 3, 3, 3, 9}, {3, 3, 3, 3, 3, 4, 4, 4}, {3, 3, 3, 3, 3, 4, 8}, {3, 3, 3, 3, 3, 12}, {3, 3, 3, 3, 5, 5, 5}, {3, 3, 3, 3, 5, 10}, {3, 3, 3, 4, 4, 5, 5}, {3, 3, 3, 4, 4, 10}, {3, 3, 3, 4, 5, 9}, {3, 3, 3, 5, 5, 8}, {3, 3, 3, 8, 10}, {3, 3, 3, 9, 9}, {3, 3, 4, 4, 4, 4, 5}, {3, 3, 4, 4, 4, 9}, {3, 3, 4, 4, 5, 8}, {3, 3, 4, 5, 12}, {3, 3, 4, 8, 9}, {3, 3, 5, 8, 8}, {3, 3, 9, 12}, {3, 4, 4, 4, 4, 4, 4}, {3, 4, 4, 4, 4, 8}, {3, 4, 4, 4, 12}, {3, 4, 4, 8, 8}, {3, 4, 5, 5, 5, 5}, {3, 4, 5, 5, 10}, {3, 4, 8, 12}, {3, 4, 10, 10}, {3, 5, 5, 5, 9}, {3, 5, 9, 10}, {3, 8, 8, 8}, {3, 12, 12}, {4, 4, 4, 5, 5, 5}, {4, 4, 4, 5, 10}, {4, 4, 5, 5, 9}, {4, 4, 9, 10}, {4, 5, 5, 5, 8}, {4, 5, 8, 10}, {4, 5, 9, 9}, {5, 5, 5, 12}, {5, 5, 8, 9}, {5, 10, 12}, {8, 9, 10}, {9, 9, 9}})

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
