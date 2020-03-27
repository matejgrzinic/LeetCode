package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}

func searchInsert(nums []int, target int) int {
	for index := range nums {
		if target <= nums[index] {
			return index
		}
	}
	return len(nums)
}
