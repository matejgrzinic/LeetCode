// https://leetcode.com/problems/valid-parentheses/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 30.00% of Go online submissions for Valid Parentheses.

package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	m := map[string]int{"(": 0, "[": 0, "{": 0}
	stack := []string{}
	for _, char := range s {
		if _, ok := m[string(char)]; ok {
			m[string(char)]++
			stack = append(stack, string(char))
		} else {
			if string(char) == ")" {
				if !isValidAndUpdateStack("(", ")", m, &stack) {
					return false
				}
			} else if string(char) == "}" {
				if !isValidAndUpdateStack("{", "}", m, &stack) {
					return false
				}
			} else {
				if !isValidAndUpdateStack("[", "]", m, &stack) {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isValidAndUpdateStack(open string, close string, m map[string]int, stack *[]string) bool {

	if m[open] > 0 && len(*stack) > 0 && (*stack)[len(*stack)-1] == open {
		m[open]--
		*stack = (*stack)[:len(*stack)-1]
		return true
	}
	return false
}
