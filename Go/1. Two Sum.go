// https://leetcode.com/problems/two-sum/

// Runtime: 4 ms, faster than 94.74% of Go online submissions for Two Sum.
// Memory Usage: 3.8 MB, less than 7.69% of Go online submissions for Two Sum.

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for index, value := range nums {
		if num, ok := seen[value]; ok {
			return []int{num, index}
		} else {
			seen[target-value] = index
		}
	}

	return []int{}
}
