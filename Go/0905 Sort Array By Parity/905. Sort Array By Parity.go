package main

import "./test0905"

func main() {
	test0905.Test(sortArrayByParity)
}

func sortArrayByParity(A []int) []int {
	l, r := 0, len(A)-1
	for l < r {
		if A[l]%2 == 1 {
			for l < r && A[r]%2 == 1 {
				r--
			}
			A[l], A[r] = A[r], A[l]
			r--
		}
		l++
	}
	return A
}
