package main

import (
	"strconv"

	"./test0412"
)

func main() {
	test0412.Test(fizzBuzz)
}

func fizzBuzz(n int) []string {
	words := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			words = append(words, "FizzBuzz")
		} else if i%5 == 0 {
			words = append(words, "Buzz")
		} else if i%3 == 0 {
			words = append(words, "Fizz")
		} else {
			words = append(words, strconv.Itoa(i))
		}
	}
	return words
}
