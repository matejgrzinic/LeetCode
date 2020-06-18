package main

import (
	"strconv"
	"strings"

	"./test0038"
)

func main() {
	test0038.Test(countAndSay)
}

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		newS := []string{}
		last, c := s[0], 0
		for j := 1; j < len(s); j++ {
			c++
			if s[j] != last {
				newS = append(newS, strconv.Itoa(c), string(last))
				c = 0
				last = s[j]
			}
		}
		newS = append(newS, strconv.Itoa(c+1), string(last))
		s = strings.Join(newS, "")
	}
	return s
}
