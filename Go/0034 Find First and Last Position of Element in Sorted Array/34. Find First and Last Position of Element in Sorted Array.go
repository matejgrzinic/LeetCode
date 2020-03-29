package main

import "./test0034"

func main() {
	test0034.Test(searchRange)
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, h := binary(nums, target, 0, len(nums)-1)
	return []int{l, h}
}

func binary(nums []int, target int, left int, right int) (int, int) {
	mid := (left + right) / 2
	low, high := -1, -1

	if nums[mid] == target {
		low, high = updateLowHigh(low, high, mid, mid)
	}
	if left >= right {
		return low, high
	}
	if target <= nums[mid] {
		l, h := binary(nums, target, left, mid-1)
		low, high = updateLowHigh(low, high, l, h)
	}

	l, h := binary(nums, target, mid+1, right)
	low, high = updateLowHigh(low, high, l, h)
	return low, high
}

func updateLowHigh(low int, high int, l int, h int) (int, int) {
	if (l < low || low == -1) && l > -1 {
		low = l
	}
	if h > high {
		high = h
	}
	return low, high
}
