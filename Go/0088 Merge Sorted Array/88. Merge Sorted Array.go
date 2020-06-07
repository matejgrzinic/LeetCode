package main

import "./test0088"

func main() {
	test0088.Test(merge)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	index, n2Index := 0, 0

	for {
		if m == 0 && n == 0 {
			return
		}
		if m == 0 {
			nums1[index] = nums2[n2Index]
			n2Index++
			index++
			n--
		} else if n > 0 {
			if nums2[n2Index] < nums1[index] {
				shiftRight(nums1, index)
				nums1[index] = nums2[n2Index]
				n2Index++
				n--
			} else {
				m--
			}
			index++
		} else {
			break
		}
	}
}

func shiftRight(nums []int, index int) {
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
}
