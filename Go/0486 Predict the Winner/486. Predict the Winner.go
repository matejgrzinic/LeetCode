package main

func PredictTheWinner(nums []int) bool {
	return f(nums, 0, 0, true)
}

func f(nums []int, p1 int, p2 int, p1Turn bool) bool {
	if len(nums) == 0 {
		return p1 >= p2
	}

	if p1Turn {
		return f(nums[1:], p1+nums[0], p2, false) || f(nums[:len(nums)-1], p1+nums[len(nums)-1], p2, false)
	} else {
		return f(nums[1:], p1, p2+nums[0], true) && f(nums[:len(nums)-1], p1, p2+nums[len(nums)-1], true)
	}
}
