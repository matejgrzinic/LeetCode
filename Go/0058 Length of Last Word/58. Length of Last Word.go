package main

import (
	"./test0058"
)

func main() {
	test0058.Test(lengthOfLastWord)
}

func lengthOfLastWord(s string) int {
	indexFirst := len(s)
	seen := false
	for i := len(s) - 1; i >= 0; i-- {
		if !seen {
			if string(s[i]) == " " {
				indexFirst--
			} else {
				seen = true
			}
		}
		if seen && string(s[i]) == " " {
			return indexFirst - 1 - i
		}
	}
	return indexFirst
}
