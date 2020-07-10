package main

import (
	"./test0168"
)

func main() {
	test0168.Test(convertToTitle)
}

func convertToTitle(n int) string {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	n--
	result := ""
	for n >= 0 {
		num := n % len(letters)
		result = letters[num] + result //letters[int(n/len(letters))]
		if num == 0 {
			n = n - len(letters)
		}
		n = n - num
		if n == 0 {
			break
		}
	}
	return result
}
