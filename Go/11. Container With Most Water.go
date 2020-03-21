// https://leetcode.com/problems/container-with-most-water/

// Runtime: 12 ms, faster than 92.06% of Go online submissions for Container With Most Water.
// Memory Usage: 5.8 MB, less than 46.67% of Go online submissions for Container With Most Water.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	best := 0
	for left, right := 0, len(height)-1; left < right; {
		score := 0
		if height[left] < height[right] {
			score = height[left] * (right - left)
			left++
		} else {
			score = height[right] * (right - left)
			right--
		}
		if score > best {
			best = score
		}
	}

	return best
}
