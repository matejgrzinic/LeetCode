package main

import (
	"./test0215"
)

func main(){
	test0215.Test(findKthLargest)
}

func findKthLargest(nums []int, k int) int {
	firstParent := int((len(nums)-1)/2)
	for i := firstParent; i >= 0; i-- {			
		siftDown(&nums, i)
	}
	for i := 0; i < k-1; i++ {
		n := len(nums)
		nums[0], nums[n-1] = nums[n-1], nums[0]
		nums = nums[:n-1]
		siftDown(&nums, 0)		
	}	
	return nums[0]
}

func siftDown(nums *[]int, index int){
	leftChild := index * 2 + 1
	rightChild := index * 2 + 2

	if leftChild >= len(*nums){
		return
	}
	bigger := leftChild
	if rightChild < len(*nums) && (*nums)[rightChild] > (*nums)[leftChild]{
		bigger = rightChild
	}			

	if (*nums)[bigger] > (*nums)[index] {
		(*nums)[bigger], (*nums)[index] = (*nums)[index], (*nums)[bigger]
		siftDown(nums, bigger)
	} else {
		return
	}
}