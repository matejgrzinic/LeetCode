package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum, result := 0, len(nums)+1

	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			result = myMin(result, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}

	if result == len(nums)+1 {
		return 0
	}
	return result
}

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
