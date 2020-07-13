package main

import (
	"strconv"
	"strings"

	"./test0043"
)

func main() {
	test0043.Test(multiply)
}

func multiply(num1 string, num2 string) string {
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	results := make([]string, len(num2))
	index1 := 1

	for i := len(num2) - 1; i >= 0; i-- {
		result := make([]string, len(num1))
		num2, _ := strconv.Atoi(string(num2[i]))
		index2 := 1
		carry := 0
		for j := len(num1) - 1; j >= 0; j-- {
			num1, _ := strconv.Atoi(string(num1[j]))
			product := num1*num2 + carry
			carry = int(product / 10)
			product = product % 10
			result[len(result)-index2] = strconv.Itoa(product)
			index2++
		}
		if carry > 0 {
			result = append([]string{strconv.Itoa(carry)}, result...)
		}
		for k := 0; k < index1-1; k++ {
			result = append(result, "0")
		}
		results[len(results)-index1] = strings.Join(result, "")
		index1++
	}

	finalString := results[0]
	for i := 1; i < len(results); i++ {
		finalString = stringAdd(finalString, results[i])
	}
	return finalString
}

func stringAdd(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	result := make([]string, len(a))
	index := 1
	carry := false

	for i := len(b) - 1; i >= 0; i-- {
		num1, _ := strconv.Atoi(string(b[i]))
		num2, _ := strconv.Atoi(string(a[len(a)-index]))
		sum := num1 + num2

		if carry {
			sum++
			carry = false
		}

		if sum > 9 {
			sum -= 10
			carry = true
		}

		result[len(a)-index] = strconv.Itoa(sum)
		index++
	}

	for i := len(a) - index; i >= 0; i-- {
		sum, _ := strconv.Atoi(string(a[i]))
		if carry {
			sum++
			carry = false
		}

		if sum > 9 {
			sum -= 10
			carry = true
		}
		result[i] = strconv.Itoa(sum)
	}
	if carry {
		result = append([]string{"1"}, result...)
	}
	return strings.Join(result, "")
}
