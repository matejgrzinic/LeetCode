package main

import (
	"strconv"

	"./test0179"
)

func main() {
	test0179.Test(largestNumber)
}

func largestNumber(nums []int) string {
	result := ""
	for len(nums) > 0 {
		index := -1
		max := 0
		for i := 0; i < len(nums); i++ {
			if compare(nums[i], max) {
				max = nums[i]
				index = i
			}
		}
		nums[index], nums[len(nums)-1] = nums[len(nums)-1], nums[index]
		nums = nums[:len(nums)-1]
		result += strconv.Itoa(max)
	}

	if len(result) == 0 || result[0] == '0' {
		return "0"
	}
	//return strings.Join(stringValues, "")
	return result
}

func compare(a int, b int) bool {
	aS := strconv.Itoa(a)
	bS := strconv.Itoa(b)

	n := len(aS)
	m := len(bS)

	for i := 0; true; i++ {
		sA, sB := string(aS[n-1]), string(bS[m-1])

		if i >= n && i >= m {
			return true
		}

		if i < n {
			sA = string(aS[i])
		}

		if i < m {
			sB = string(bS[i])
		}

		if sA > sB {
			return true
		} else if sA < sB {
			return false
		}
	}
	return true
}
