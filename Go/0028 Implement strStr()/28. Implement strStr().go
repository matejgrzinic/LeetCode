// https://leetcode.com/problems/implement-strstr/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Implement strStr().

package main

import "fmt"

func main() {
	fmt.Println(strStr("a", "a"))            // 0
	fmt.Println(strStr("mississippi", "pi")) // 9
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || haystack == needle {
		return 0
	}

	for chIndex := 0; chIndex <= len(haystack)-len(needle); chIndex++ {
		found := true
		for eqIndex := 0; eqIndex < len(needle); eqIndex++ {
			if needle[eqIndex] != haystack[chIndex+eqIndex] {
				found = false
				break
			}
		}
		if found {
			return chIndex
		}

	}
	return -1
}
