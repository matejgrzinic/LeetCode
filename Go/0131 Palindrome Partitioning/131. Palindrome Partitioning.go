package main

import (
	"./test0131"
)

func main() {
	test0131.Test(partition)
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	result := [][]string{}
	helper(s, []string{}, &result)
	return result
}

func helper(s string, current []string, result *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
		return
	}
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			current = append(current, s[:i])
			helper(s[i:], current, result)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
