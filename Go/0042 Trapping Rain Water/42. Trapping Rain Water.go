package main

import (
	"./test0042"
)

func main() {
	test0042.Test(trap)
}

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	l, r := 0, len(height)-1
	lmax, rmax := height[l], height[r]
	sum := 0
	leftT := false
	if lmax < rmax {
		leftT = true
	}
	for l < r {
		s := min(lmax, rmax)
		if leftT {
			s -= height[l]
		} else {
			s -= height[r]
		}
		if s > 0 {
			sum += s
		}
		if lmax < rmax {
			l++
			lmax = max(lmax, height[l])
			leftT = true
		} else {
			r--
			rmax = max(rmax, height[r])
			leftT = false
		}
	}
	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
