package main

import (
	"./test0050"
)

func main() {
	test0050.Test(myPow)
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = n * -1
	}

	k := 1.0
	for n != 1 {
		if n%2 == 0 {
			x *= x
			n /= 2
		} else {
			k *= x
			n--
		}
	}
	x *= k
	return float64(int(x*100000)) / 100000
}
