package main

import "./test0169"

func main() {
	test0169.Test(majorityElement)
}

func majorityElement(nums []int) int {
	if len(nums) < 3 {
		return nums[0]
	}
	counter := map[int]int{}
	majority := int((len(nums) + 1) / 2)
	for _, num := range nums {
		if val, ok := counter[num]; ok {
			val++
			if val == majority {
				return num
			}
			counter[num] = val
		} else {
			counter[num] = 1
		}
	}
	return majority
}

func majorityElement2(nums []int) int {
	if len(nums) < 3 {
		return nums[0]
	}
	n := len(nums)
	majority := int((len(nums) + 1) / 2)

	for {
		num := nums[0]
		if n == majority {
			return num
		}
		nums[0], nums[n-1] = nums[n-1], nums[0]
		n--
		counter := 1

		for index := 0; index < n; index++ {
			if nums[index] == num {
				nums[index], nums[n-1] = nums[n-1], nums[index]
				n--
				counter++
				index--
			}
		}
		if counter >= majority {
			return num
		}
	}
}
