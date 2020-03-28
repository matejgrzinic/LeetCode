package main

import "fmt"

func main() {
	test()
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	mid := (right + left) / 2
	pivot := nums[left]

	if nums[mid] == target {
		return mid
	}
	if left >= right {
		return -1
	}

	if target > nums[mid] {
		if target >= pivot && nums[mid] < pivot {
			return binarySearch(nums, left, mid-1, target)
		}
		return binarySearch(nums, mid+1, right, target)
	}
	if nums[mid] < pivot || target >= pivot {
		return binarySearch(nums, left, mid-1, target)
	}
	return binarySearch(nums, mid+1, right, target)
}

// Start of tests

var numTests int
var failed int

func test() {
	numTests = 0
	failed = 0
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 4, 0)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 5, 1)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 6, 2)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 7, 3)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 0, 4)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 1, 5)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 2, 6)
	equals([]int{4, 5, 6, 7, 0, 1, 2}, 3, -1)
	equals([]int{5, 1, 3}, 2, -1)
	equals([]int{8, 1, 3}, 8, 0)
	equals([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4)
	equals([]int{3, 1}, 1, 1)
	equals([]int{5, 1, 3}, 3, 2)
	equals([]int{1, 3}, 3, 1)
	equals([]int{5, 1, 2, 3, 4}, 1, 1)

	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
}

func equals(nums []int, target int, expected int) {
	numTests++
	got := search(nums, target)
	if got != expected {
		fmt.Println(nums, target, "got:", got, "expected:", expected)
		failed++
	}
}
