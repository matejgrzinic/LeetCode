package main

import "./test0066"

func main() {
	test0066.Test(plusOne)
}

func plusOne(digits []int) []int {
	n := len(digits) - 1

	// Creates copy of digits into digits so tests can show initial digits values
	t := make([]int, n+1)
	copy(t, digits)
	digits = t

	for i := 0; i <= n; i++ {
		if digits[n-i] == 9 {
			digits[n-i] = 0
		} else {
			digits[n-i]++
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}
