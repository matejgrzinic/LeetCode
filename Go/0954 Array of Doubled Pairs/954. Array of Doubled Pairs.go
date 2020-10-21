package main

import (
	"sort"
)

func canReorderDoubled(A []int) bool {
	myNums := make(map[int]int)
	for _, num := range A {
		myNums[num]++
	}
	sort.Slice(A, func(i int, j int) bool {
		a, b := A[i], A[j]
		if a < 0 && b < 0 {
			return a > b
		}
		return a < b
	})
	for _, num := range A {
		if myNums[num] == 0 {
			continue
		}
		if myNums[num*2] == 0 {
			return false
		}
		myNums[num*2]--
		myNums[num]--
	}
	return true
}
