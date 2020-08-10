package main

import "./test0896"

func main() {
	test0896.Test(isMonotonic)
}

func isMonotonic(A []int) bool {
	if len(A) < 3 {
		return true
	}
	mode := 0
	for i := 1; i < len(A); i++ {
		if mode == 0 {
			if A[i] > A[i-1] {
				mode = 1
			} else if A[i] < A[i-1] {
				mode = -1
			}
		} else if mode == 1 && A[i] < A[i-1] {
			return false
		} else if mode == -1 && A[i] > A[i-1] {
			return false
		}
	}
	return true
}
