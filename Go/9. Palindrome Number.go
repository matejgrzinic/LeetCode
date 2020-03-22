// https://leetcode.com/problems/palindrome-number/

// Runtime: 4 ms, faster than 98.99% of Go online submissions for Palindrome Number.
// Memory Usage: 5.4 MB, less than 25.00% of Go online submissions for Palindrome Number.

package main

import (
	"strconv"
)

// func main() {
// 	fmt.Println(isPalindrome(121))
// }

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	n := len(str) - 1
	for i := 0; i < n; i++ {
		if str[i] != str[n-i] {
			return false
		}
	}

	return true
}
