package main

import (
	"./test0091"
)

func main() {
	test0091.Test(numDecodings)
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	x, y := 1, 1
	if s[0] == '0' {
		y = 0
	}

	for i := 1; i < len(s); i++ {
		t := 0
		if s[i] != '0' {
			t += y
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			t += x
		}
		x, y = y, t
	}
	return y
}
