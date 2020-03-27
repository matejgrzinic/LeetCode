// Runtime: 0 ms, faster than 100.00% of Go online submissions for Next Permutation.
// Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Next Permutation.

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	x := -1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			x = i - 1
		}
	}
	if x == -1 {
		sort.Ints(nums)
	} else {
		y := -1
		for i := 0; i < len(nums); i++ {
			if nums[x] < nums[i] {
				y = i
			}
		}
		nums[x], nums[y] = nums[y], nums[x]

		for i, j := x+1, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
