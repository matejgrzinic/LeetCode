// https://leetcode.com/problems/4sum/

// Runtime: 8 ms, faster than 89.43% of Go online submissions for 4Sum.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for 4Sum.

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	// var results [][]int
	results := [][]int{}
	sort.Ints(nums)

	for a := 0; a < len(nums); a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums); b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for c, d := b+1, len(nums)-1; c < d; {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					results = append(results, []int{nums[a], nums[b], nums[c], nums[d]})

					d--
					for c < d && nums[d] == nums[d+1] {
						d--
					}

					c++
					for c < d && nums[c] == nums[c-1] {
						c++
					}

				} else if sum > target {
					d--
				} else {
					c++
				}
			}
		}
	}

	return results
}
