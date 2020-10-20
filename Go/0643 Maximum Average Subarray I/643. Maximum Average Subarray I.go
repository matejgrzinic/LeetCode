package main

func findMaxAverage(nums []int, k int) float64 {
	sum, max := 0, 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max = sum
	for l := 0; l+k < len(nums); l++ {
		sum = sum - nums[l] + nums[l+k]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
