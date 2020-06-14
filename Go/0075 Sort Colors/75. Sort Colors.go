package main

import "./test0075"

func main() {
	test0075.Test(sortColors)
}

func sortColors(nums []int) {
	l, m, r := 0, 0, len(nums)-1

	for m <= r {
		if nums[m] == 0 {
			nums[l], nums[m] = nums[m], nums[l]
			m++
			l++
		} else if nums[m] == 1 {
			m++
		} else if nums[m] == 2 {
			nums[r], nums[m] = nums[m], nums[r]
			r--
		}
	}
}
