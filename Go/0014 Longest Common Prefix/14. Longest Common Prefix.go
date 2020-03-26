// https://leetcode.com/problems/longest-common-prefix/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
// Memory Usage: 2.5 MB, less than 25.00% of Go online submissions for Longest Common Prefix.

package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix(([]string{"flower", "flow", "flight"})))
}

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}
	for index := 0; true; index++ {
		var char byte
		for i, word := range strs {
			if i == 0 && len(word) > index {
				char = word[index]
			} else if len(word) <= index || word[index] != char {
				return result
			}
		}
		result += string(char)
	}
	return result
}
