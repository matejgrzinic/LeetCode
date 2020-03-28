package main

import "./test0033"

func main() {
	test0033.Test(search)
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
