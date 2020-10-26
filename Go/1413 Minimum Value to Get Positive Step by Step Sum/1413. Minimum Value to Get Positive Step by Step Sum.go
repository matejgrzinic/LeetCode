package main

func minStartValue(nums []int) int {
	sum, min := nums[0], nums[0]
	for _, num := range nums[1:] {
		sum += num
		if sum < min {
			min = sum
		}
	}
	min = -min + 1
	if min < 1 {
		min = 1
	}
	return min
}
