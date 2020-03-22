// https://leetcode.com/problems/integer-to-roman/

// Runtime: 4 ms, faster than 94.25% of Go online submissions for Integer to Roman.
// Memory Usage: 3.5 MB, less than 58.33% of Go online submissions for Integer to Roman.

package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	symbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	str := ""

	for curr := 12; curr >= 0; {
		if numbers[curr] > num {
			curr--
		} else {
			str += symbols[curr]
			num -= numbers[curr]
			if num == 0 {
				return str
			}
		}
	}

	return str
}
