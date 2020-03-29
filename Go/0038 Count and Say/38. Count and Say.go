package main

import "./test0038"

func main() {
	test0038.Test(countAndSay)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	a := "2"
	return "1" + a
}
