package main

import (
	"./test0137"
)

func main() {
	test0137.Test(singleNumber)
}

func singleNumber(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num] = true
		} else {
			m[num] = false
		}
	}
	for k, v := range m {
		if v == false {
			return k
		}
	}
	return 0
}

// https://youtu.be/cOFAmaMBVps?t=484
func singleNumberBitManipulation(nums []int) int {
	once := 0
	twice := 0
	for _, num := range nums {
		once = ^twice & (once ^ num)
		twice = ^once & (twice ^ num)
	}
	return once
}
