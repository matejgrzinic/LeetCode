// https://leetcode.com/problems/3sum-closest/submissions/

// Runtime: 4 ms, faster than 96.48% of Go online submissions for 3Sum Closest.
// Memory Usage: 2.7 MB, less than 50.00% of Go online submissions for 3Sum Closest.

package main

import (
	"math"
	"sort"
)

// func main() {
// 	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))    // 2
// 	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100))    // 3
// 	fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, 3)) // 3
// }

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	best := math.MaxInt32
	bestsum := math.MaxInt32

	for index, num := range nums {
		l := index + 1
		r := len(nums) - 1

		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		if index > len(nums)-3 {
			return bestsum
		}

		for l < r {
			sum := num + nums[l] + nums[r]
			diff := intAbs(sum - target)
			if diff < best {
				best = diff
				bestsum = sum
			} else if sum > target {
				r--
			} else {
				l++
			}
		}

	}
	return bestsum
}

func intAbs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}
