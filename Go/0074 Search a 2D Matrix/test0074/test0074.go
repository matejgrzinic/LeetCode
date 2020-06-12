package test0074

import (
	"fmt"
	"time"
)

type testFunc func([][]int, int) bool

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn

	test1 := createNxMMatrix(15000, 1500)
	start := time.Now()

	equals([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 1, true)
	equals([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 7, true)
	equals([][]int{{-10}, {-7}, {-5}}, -10, true)
	equals([][]int{{1, 3}}, 1, true)
	equals([][]int{{1}, {3}}, 3, true)
	equals([][]int{{1, 1}}, 2, false)
	equals([][]int{{1}}, 1, true)
	equals([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3, true)
	equals([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13, false)
	equals(test1, 15000*1500-1, true)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(matrix [][]int, target int, expected bool) {
	numTests++
	got := testFunction(matrix, target)
	if got != expected {
		fmt.Println(matrix, target, "got:", got, "expected:", expected)
		failed++
	}
}

func createNxMMatrix(n int, m int) [][]int {
	c := 1
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, m)
		for j := 0; j < m; j++ {
			r[i][j] = c
			c++
		}
	}
	return r
}
