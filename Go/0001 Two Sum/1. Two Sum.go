package main

import "./test0001"

func main() {
	test0001.Test(twoSum)
}

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for index, value := range nums {
		if num, ok := seen[value]; ok {
			return []int{num, index}
		}
		seen[target-value] = index
	}

	return []int{}
}
