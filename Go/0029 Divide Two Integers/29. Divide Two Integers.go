// https://leetcode.com/problems/divide-two-integers/

// Runtime: 2460 ms, faster than 31.12% of Go online submissions for Divide Two Integers.
// Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Divide Two Integers.

package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))           // 3
	fmt.Println(divide(7, -3))           // -2
	fmt.Println(divide(-2147483648, -1)) // -2
}

func divide(dividend int, divisor int) int {
	negative := false
	if dividend < 0 {
		dividend *= -1
		negative = true
	}
	if divisor < 0 {
		divisor *= -1
		negative = !negative
	}
	result := 0
	sum := 0
	for sum <= dividend {
		sum += divisor
		result++
	}

	if negative {
		return -result + 1
	}
	if result < 2147483647 {
		return result - 1
	}
	return 2147483647

}
