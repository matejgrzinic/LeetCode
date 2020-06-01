package main

import (
	"./test0067"
)

func main() {
	test0067.Test(addBinary)
}

func addBinary(a string, b string) string {
	n, m := len(a)-1, len(b)-1

	if n < m {
		a, b = b, a
		n, m = m, n
	}

	result := ""

	c := false
	for i := 0; i <= n; i++ {
		sum := 0
		if a[n-i] == '1' {
			sum++
		}
		if i <= m {
			if b[m-i] == '1' {
				sum++
			}
		}
		if c {
			sum++
			c = false
		}
		if sum > 1 {
			c = true
			sum = sum % 2
		}

		ch := "0"
		if sum == 1 {
			ch = "1"
		}
		result = ch + result
	}
	if c {
		result = "1" + result
	}

	return result
}
