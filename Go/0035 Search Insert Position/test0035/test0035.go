package test0035

import "fmt"

type testFunc func([]int, int) int

var numTests int
var failed int
var testFunction testFunc

func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn

	equals([]int{1, 3, 5, 6}, 5, 2)
	equals([]int{1, 2}, 1, 0)

	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
}

func equals(nums []int, target int, expected int) {
	numTests++
	got := testFunction(nums, target)
	if got != expected {
		fmt.Println(nums, target, "got:", got, "expected:", expected)
		failed++
	}
}
