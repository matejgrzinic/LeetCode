// https://leetcode.com/problems/generate-parentheses

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Generate Parentheses.
// Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Generate Parentheses.

package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	result := []string{}
	generateParenthesisStep("", n, 0, &result)
	return result
}

func generateParenthesisStep(current string, toOpen int, toClose int, result *[]string) {
	if toOpen == 0 {
		if toClose == 0 {
			*result = append(*result, current)
			return
		} else {
			generateParenthesisStep(current+")", 0, toClose-1, result)
		}
	} else {
		generateParenthesisStep(current+"(", toOpen-1, toClose+1, result)
		if toClose > 0 {
			generateParenthesisStep(current+")", toOpen, toClose-1, result)
		}
	}
}
