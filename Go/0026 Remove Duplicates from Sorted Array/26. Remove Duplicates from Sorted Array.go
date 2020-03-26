// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// Runtime: 8 ms, faster than 90.64% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.6 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.

package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	unique := 0
	for index := range nums[1:] {
		if nums[index+1] != nums[unique] {
			unique++
			nums[unique] = nums[index+1]
		}
	}
	return unique + 1
}
