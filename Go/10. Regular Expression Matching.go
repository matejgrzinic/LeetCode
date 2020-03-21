// https://leetcode.com/problems/regular-expression-matching/submissions/

// Runtime: 4 ms, faster than 56.28% of Go online submissions for Regular Expression Matching.
// Memory Usage: 3.3 MB, less than 20.00% of Go online submissions for Regular Expression Matching.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*ip*."))
}

var (
	m = make(map[string]bool)
)

func isMatch(s string, p string) bool {

	if _, ok := m[s+"_"+p]; ok {
		return false
	} else {
		m[s+"_"+p] = true
	}

	if len(s) == 0 {
		for len(p) > 1 && p[1] == '*' {
			p = p[2:]
		}
		if len(p) == 0 {
			return true
		}

		return false
	}

	if len(p) == 0 {
		return false
	}

	if len(p) > 1 && p[1] == '*' {
		if s[0] == p[0] || p[0] == '.' {
			if isMatch(s[1:], p) {
				return true
			}
		}
		if isMatch(s, p[2:]) {
			return true
		}
	} else {
		if s[0] == p[0] || p[0] == '.' {
			if isMatch(s[1:], p[1:]) {
				return true
			}
		}
	}
	return false
}
