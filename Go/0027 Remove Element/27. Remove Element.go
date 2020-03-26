// https://leetcode.com/problems/remove-element

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Remove Element.

package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(nums)
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)

	fmt.Println()

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(nums)
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	newIndex := 0
	for index := range nums {
		if nums[index] != val {
			nums[newIndex] = nums[index]
			newIndex++
		}
	}
	return newIndex
}
