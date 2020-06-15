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
	return float64(int(calcPow(x, n)*100000)) / 100000
}

func calcPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		return calcPow(x*x, n/2)
	}
	return calcPow(x, n-1) * x
}
