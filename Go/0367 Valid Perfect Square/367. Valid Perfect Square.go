package main

import (
	"./test0367"
)

func main() {
	test0367.Test(isPerfectSquare)
}

func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
