package main

import (
	"strconv"
	"strings"

	"./test0093"
)

func main() {
	test0093.Test(restoreIpAddresses)
}

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}
	nums := []int{}
	for _, ch := range s {
		num, _ := strconv.Atoi(string(ch))
		nums = append(nums, num)
	}
	ints := helper(nums[1:], [][]int{{nums[0]}})
	results := make([]string, len(ints))
	for i, ip := range ints {
		tmp := make([]string, 4)
		for j, part := range ip {
			tmp[j] = strconv.Itoa(part)
		}
		results[i] = strings.Join(tmp, ".")
	}
	return results
}

func helper(nums []int, results [][]int) [][]int {
	if len(nums) == 0 {
		newResults := [][]int{}
		for _, ip := range results {
			if len(ip) == 4 {
				newResults = append(newResults, ip)
			}
		}
		return newResults
	}
	num := nums[0]
	newResults := [][]int{}

	for _, ip := range results {
		l := len(ip) - 1

		max := 25
		if num > 5 {
			max = 24
		}

		if ip[l] <= max && ip[l] != 0 {
			tmp := make([]int, len(ip))
			copy(tmp, ip)
			tmp[l] = tmp[l]*10 + num
			newResults = append(newResults, tmp)
		}

		if l < 3 && (3-l)*3 >= len(nums) {
			tmp := make([]int, len(ip)+1)
			copy(tmp, ip)
			tmp[l+1] = num
			newResults = append(newResults, tmp)
		}
	}
	return helper(nums[1:], newResults)
}
