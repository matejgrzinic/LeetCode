package main

import "./test0053"

func main() {
	test0053.Test(maxSubArray)
}

func maxSubArray(nums []int) int {
	currentSum, best := 0, nums[0]
	for _, num := range nums {
		currentSum += num
		if currentSum > num {
		} else {
			currentSum = num
		}
		if currentSum > best {
			best = currentSum
		}
	}
	return best
}
