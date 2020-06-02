package main

import "./test0069"

func main() {
	test0069.Test(mySqrt)
}

func mySqrt(x int) int {
	i := 0
	for {
		if i*i > x {
			return i - 1
		}
		i++
	}
	return 0
}
