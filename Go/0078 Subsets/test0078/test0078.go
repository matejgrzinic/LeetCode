package test0078

import (
	"fmt"
	"time"
)

type testFunc func([]int) [][]int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{1, 2, 3}, [][]int{{3}, {1}, {2}, {1, 2, 3}, {1, 2}, {2, 3}, {1, 3}, {}})
	equals([]int{0}, [][]int{{0}, {}})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, expected [][]int) {
	numTests++
	got := testFunction(nums)
	if !sliceEqualsDifferentOrder(got, expected) {
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

func duplicate(matrix [][]int) [][]int {
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func sliceEqualsDifferentOrder(slice1 [][]int, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	s2dup := duplicate(slice2)

	for i := 0; i < len(slice1); i++ {
		found := false
		for k := 0; k < len(s2dup); k++ {
			if len(slice1[i]) == len(s2dup[k]) {
				found = true
				for j := 0; j < len(s2dup[k]); j++ {
					if slice1[i][j] != s2dup[k][j] {
						found = false
						break
					}
				}
				if found {
					s2dup[k], s2dup[len(s2dup)-1] = s2dup[len(s2dup)-1], s2dup[k]
					s2dup = s2dup[:len(s2dup)-1]
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
