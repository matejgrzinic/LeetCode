package main

func findTargetSumWays(nums []int, S int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum < S || (totalSum+S)%2 != 0 {
		return 0
	}

	target := (totalSum + S) / 2
	dp := make([]int, target+1)
	dp[0] = 1

	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}
