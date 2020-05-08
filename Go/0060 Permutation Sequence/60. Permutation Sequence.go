package main

import (
	"strconv"

	"./test0060"
)

func main() {
	test0060.Test(getPermutation)
}

func getPermutation(n int, k int) string {
	elements := []int{}
	for i := 0; i < n; i++ {
		elements = append(elements, i+1)
	}

	for i := 0; i < k-1; i++ {
		nextPermutation(elements)
	}

	return sliceToString(elements)
}

func nextPermutation(elements []int) {
	x := -1
	for i := 0; i < len(elements)-1; i++ {
		if elements[i] < elements[i+1] {
			x = i
		}
	}

	if x == -1 {
		return
	}

	y := -1
	for i := 0; i < len(elements); i++ {
		if elements[x] < elements[i] {
			y = i
		}
	}

	if y == -1 {
		return
	}

	elements[x], elements[y] = elements[y], elements[x]

	for i := 0; x+1+i < len(elements)-i-1; i++ {
		elements[x+1+i], elements[len(elements)-i-1] = elements[len(elements)-i-1], elements[x+1+i]
	}
}

func sliceToString(elements []int) string {
	result := ""
	for i := 0; i < len(elements); i++ {
		result += strconv.Itoa(elements[i])
	}
	return result
}
