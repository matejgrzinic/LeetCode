package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func lastStoneWeightII(stones []int) int {
	set := map[int]struct{}{0: {}}

	for _, stone := range stones {
		set2 := set
		set = map[int]struct{}{} // clear set
		for tmpResult := range set2 {
			set[tmpResult+stone] = struct{}{}
			set[tmpResult-stone] = struct{}{}
		}
	}
	m := math.MaxInt32
	for k := range set {
		if k >= 0 && k < m {
			m = k
		}
	}
	return m
}
