package main

import (
	"unicode"

	"./test0125"
)

func main() {
	test0125.Test(isPalindrome)
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < len(s) && !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
		}
		for r >= 0 && !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
		}

		if r <= l {
			break
		}

		a, b := unicode.ToLower(rune(s[l])), unicode.ToLower(rune(s[r]))

		if a != b {
			return false
		}

		l++
		r--
	}
	return true
}
