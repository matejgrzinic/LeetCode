package main

import (
	"./test0167"
)

func main() {
	test0167.Test(twoSum)
}

func twoSum(numbers []int, target int) []int {
	nums := make(map[int]int)
	for i, num := range numbers {
		if i2, ok := nums[target-num]; ok {
			return []int{i2 + 1, i + 1}
		}
		nums[num] = i
	}
	return []int{}
}
