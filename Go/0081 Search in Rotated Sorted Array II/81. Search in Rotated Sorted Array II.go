package main

import "./test0081"

func main() {
	test0081.Test(search)
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	return binary(&nums, 0, len(nums)-1, &target)
}

func binary(nums *[]int, left int, right int, target *int) bool {
	middle := int((left + right) / 2)

	if (*nums)[middle] == *target {
		return true
	}

	if left >= right {
		return false
	}

	if *target > (*nums)[middle] {
		if (*nums)[right] >= *target {
			return binary(nums, middle+1, right, target)
		}
		return binary(nums, left, right-1, target)
	}
	if *target >= (*nums)[left] || (*target < (*nums)[right] && *target <= (*nums)[left]) {
		return binary(nums, left, right-1, target)
	}
	return binary(nums, middle+1, right, target)
}
