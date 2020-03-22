// https://leetcode.com/problems/3sum/

// Runtime: 28 ms, faster than 96.01% of Go online submissions for 3Sum.
// Memory Usage: 6.9 MB, less than 100.00% of Go online submissions for 3Sum.

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for index, num := range nums {
		l := index + 1
		r := len(nums) - 1

		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		for l < r {
			sum := num + nums[l] + nums[r]

			if sum == 0 {
				result = append(result, []int{num, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}

			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}

	}
	return result
}
