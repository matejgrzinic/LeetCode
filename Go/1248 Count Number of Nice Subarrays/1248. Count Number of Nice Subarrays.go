package main

func numberOfSubarrays(nums []int, k int) int {
	f := func(k int) int {
		result, oddCounter := 0, 0
		l, r := 0, 0
		for r < len(nums) {
			oddCounter += nums[r] % 2
			for oddCounter > k {
				oddCounter -= nums[l] % 2
				l++
			}
			result += r - l + 1
			r++
		}
		return result
	}
	return f(k) - f(k-1)
}
