package test0088

import (
	"fmt"
	"time"
)

type testFunc func([]int, int, []int, int)
var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3, []int{1,2,2,3,5,6})

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums1 []int, m int, nums2 []int, n int, expected []int) {
	numTests++
	dup := duplicate(nums1)
	testFunction(nums1, m, nums2, n)
	if !sliceEquals(nums1, expected) {
		fmt.Println(dup, nums2, "got:", nums1, "expected:", expected)
		failed++
	}
}

func duplicate(matrix []int) []int{
	duplicate := make([]int, len(matrix))
	copy(duplicate, matrix)
	return duplicate
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