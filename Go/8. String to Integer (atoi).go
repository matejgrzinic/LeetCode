// https://leetcode.com/problems/string-to-integer-atoi/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
// Memory Usage: 2.3 MB, less than 25.00% of Go online submissions for String to Integer (atoi).

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(myAtoi("-91283472332"))
}

func myAtoi(str string) int {
	trimmedString := strings.TrimSpace(str)
	num := ""
	for index, char := range trimmedString {
		ch := rune(char)
		if index == 0 && (ch == '+' || ch == '-') {
			num += string(char)
		} else if unicode.IsDigit(ch) {
			num += string(char)
		} else {
			break
		}
	}

	if len(num) == 0 || num == "+" || num == "-" {
		return 0
	} else {
		number, _ := strconv.Atoi(num)

		if number < -2147483648 {
			return -2147483648
		} else if number > 2147483647 {
			return 2147483647
		}
		return number
	}
}
