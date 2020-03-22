// https://leetcode.com/problems/roman-to-integer/

// Runtime: 4 ms, faster than 85.96% of Go online submissions for Roman to Integer.
// Memory Usage: 3.1 MB, less than 14.29% of Go online submissions for Roman to Integer.

package main

// func main() {
// 	fmt.Println(romanToInt("DCXXI"))
// }

func romanToInt(s string) int {
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	symbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	num := 0
	n := len(s)

	for curr, char := 12, 0; char < n; {
		if s[char:char+1] == symbols[curr] {
			char++
			num += numbers[curr]
		} else if char < n-1 && s[char:char+2] == symbols[curr] {
			char += 2
			num += numbers[curr]
		} else {
			curr--
		}
	}
	return num
}
